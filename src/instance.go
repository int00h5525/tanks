package main

import (
	//"fmt"
	//"strings"
	"time"
	"log"
	"net/http"
)

type Instance struct {
	Base		string
	TanksDir	string
	AttemptInterval	time.Duration
	mux		*http.ServeMux
}

func (ctx *Instance) Initialize() error {
	log.Print("Initializing context")

	ctx.mux = http.NewServeMux()

	return nil
}
