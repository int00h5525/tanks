package main

import (
	"log"
	"time"
)

func (ctx *Instance) Maintenance(maintenanceInterval time.Duration) {
	for {
		select {
			case <-time.After(maintenanceInterval):
				log.Print("Running")
		}
	}
}
