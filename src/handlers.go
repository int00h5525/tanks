package main

import (
	"net/http"
	"log"
	"strings"
	"os"
	"github.com/flosch/pongo2"
)

type FurtiveResponseWriter struct {
	w		http.ResponseWriter
	statusCode	*int
}

func (w FurtiveResponseWriter) WriteHeader(statusCode int) {
        *w.statusCode = statusCode
        w.w.WriteHeader(statusCode)
}

func (w FurtiveResponseWriter) Write(buf []byte) (n int, err error) {
        n, err = w.w.Write(buf)
        return
}

func (w FurtiveResponseWriter) Header() http.Header {
        return w.w.Header()
}

func (ctx *Instance) staticHandler(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if strings.Contains(path, "..") {
		http.Error(w, "Invalid URL path", http.StatusBadRequest)
		return
	}

	if path == "/" {
		path = "/index.html"
	}

	actual_path := ctx.ThemePath(path)
	f, err := os.Open(actual_path)
	if err != nil {
		http.NotFound(w, req)
		return
	}
	defer f.Close()

	d, err := f.Stat()
	if err != nil {
		http.NotFound(w, req)
		return
	}

	if strings.HasSuffix(actual_path, ".html") {
		tmpl, err := pongo2.FromFile(actual_path)
		if err != nil {
			log.Print(err)
			http.NotFound(w, req)
			return
		}
		tmpl.ExecuteWriter(nil, w)
	} else {
		http.ServeContent(w, req, path, d.ModTime(), f)
	}
}


func (ctx *Instance) ServeHTTP(wOrig http.ResponseWriter, r *http.Request) {
	w := FurtiveResponseWriter{
		w: 	wOrig,
		statusCode: new(int),
	}
	ctx.mux.ServeHTTP(w, r)
	log.Printf(
		"%s %s %s %d\n",
		r.RemoteAddr,
		r.Method,
		r.URL,
		*w.statusCode,
	)
}

func (ctx *Instance) BindHandlers() {
	ctx.mux.HandleFunc(ctx.Base+"/", ctx.staticHandler)
}
