package main

import (
	"fmt"
	"net/http"
	"log"
	"strings"
	"strconv"
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

func (ctx *Instance) uploadHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		http.Error(w, "Unpermitted method", http.StatusMethodNotAllowed)
		return
	}

	token := strings.TrimSpace( req.FormValue("token") )
	if token == "" {
		http.Error(w, "Invalid token", http.StatusUnprocessableEntity)
		return
	} else {
		// Validate token
	}

	newtank := Tank{}
	newtank.name = req.FormValue("name")
	newtank.author = req.FormValue("author")
	newtank.color = req.FormValue("color")
	newtank.program = req.FormValue("program")

	for i, _ := range newtank.sensors {
		var err error

		radius_raw := req.FormValue(fmt.Sprintf("s%dr", i) )
		if radius_raw == "" {
			newtank.sensors[i].radius = 0
		} else {
			newtank.sensors[i].radius, err = strconv.Atoi( radius_raw )
			if err != nil {
				newtank.sensors[i].radius = 0
			}
		}

		angle_raw := req.FormValue(fmt.Sprintf("s%da", i) )
		if angle_raw == "" {
			newtank.sensors[i].angle = 0
		} else {
			newtank.sensors[i].angle, err = strconv.Atoi( angle_raw )
			if err != nil {
				newtank.sensors[i].angle = 0
			}
		}

		width_raw := req.FormValue(fmt.Sprintf("s%dw", i) )
		if width_raw == "" {
			newtank.sensors[i].width = 0
		} else {
			newtank.sensors[i].width, err = strconv.Atoi( width_raw )
			if err != nil {
				newtank.sensors[i].width = 0
			}
		}

		turret_raw := req.FormValue( fmt.Sprintf("s%dt", i) )
		if turret_raw == "on" {
			newtank.sensors[i].is_turret = true
		} else {
			newtank.sensors[i].is_turret = false
		}
	}

	name := req.FormValue("name")
	log.Printf("Name: %s", name)
	log.Printf("Tank %v", newtank)
	ctx.saveTank(newtank, token)
	http.NotFound(w, req)
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
	ctx.mux.HandleFunc(ctx.Base+"/upload", ctx.uploadHandler)
}
