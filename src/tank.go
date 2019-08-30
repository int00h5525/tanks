package main

import (
	//"fmt"
	//"strings"
	//"time"
	//"os"
	//"path"
	"encoding/json"
	"errors"
	"log"
)

type Sensor struct {
	radius		int
	angle		int
	width		int
	is_turret	bool
}

type Tank struct {
	name		string
	author		string
	color		string
	program		string
	sensors		[10]Sensor
}

func (sensor *Sensor) validate() error {
	if sensor.radius < 0 || sensor.radius > 100 {
		return errors.New("Sensor range must be in the range 0-100")
	}

	if sensor.angle < -359 || sensor.angle > 359 {
		return errors.New("Sensor angle must be in the range -359 - 359")
	}

	if sensor.width < 0 || sensor.width > 100 {
		return errors.New("Sensor width must be in the range 0-100")
	}

	// TODO
	// Validate that program compiles

	return nil
}

func (sensor *Sensor) jsonify() []byte {
	s1, _ := json.Marshal(sensor)
	log.Printf("%s", s1)
	return s1
}

func (tank *Tank) jsonify() []byte {
	s1, _ := json.Marshal(tank)
	return s1
}

/*
func (ctx *Instance) Initialize() error {
	log.Print("Initializing context")

	ctx.mux = http.NewServeMux()

	ctx.Base = strings.TrimRight(ctx.Base, "/")

	return nil
}
*/
