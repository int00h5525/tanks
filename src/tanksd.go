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

	flag.StringVar(
		&ctx.ThemeDir,
		"theme",
		"/theme",
		"Path to static theme resources (HTML, images, css, ...)",
	)

	listen := flag.String(
		"listen",
		":8080",
		"[host]:port to bind and listen",
	)

	maintenanceInterval := flag.Duration(
		"maint",
		20*time.Second,
		"Time between rounds",
	)

	flag.Parse()

	if err := setup(); err != nil {
		log.Fatal(err)
	}

	err := ctx.Initialize()
	if err != nil {
		log.Fatal(err)
	}

	myS := &Sensor{radius:1000, angle:1000, width:10, is_turret:false}
	//myT := &Tank{name: "", author: "", color: "", program: "", sensors: [10]Sensor{*myS}}

	/*
	log.Printf("%s", *myT)
	log.Printf("%s", myT.jsonify())

	if err := myS.validate(); err != nil {
		log.Fatal(err)
	}*/

	go ctx.Maintenance(*maintenanceInterval)

	log.Printf("Sensor radius is %d", myS.radius)

	log.Printf("Listening on %s", *listen)
	log.Fatal(http.ListenAndServe(*listen, ctx))
}
