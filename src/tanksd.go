package main

import (
	"flag"
	"log"
	"math/rand"
	//"mime"
	"net/http"
	"time"
)

func setup() error {
	rand.Seed(time.Now().UnixNano())
	return nil
}

func main() {
	ctx := &Instance{}

	flag.StringVar(
		&ctx.Base,
		"base",
		"/",
		"Base URL of this instance",
	)

	listen := flag.String(
		"listen",
		":8080",
		"[host]:port to bind and listen",
	)

	flag.Parse()

	if err := setup(); err != nil {
		log.Fatal(err)
	}

	err := ctx.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	//go ctx.Maintenance(*maintenanceInterval)

	log.Printf("Listening on %s", *listen)
	log.Fatal(http.ListenAndServe(*listen, ctx))
}
