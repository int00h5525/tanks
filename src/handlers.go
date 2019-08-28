package main

import (
	"net/http"
	"log"
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
