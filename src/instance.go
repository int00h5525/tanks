package main

import (
	"fmt"
	"strings"
	"time"
	"log"
	"net/http"
	"os"
	"errors"
	"regexp"
	"path"
)

type Instance struct {
	Base		string
	StateDir	string
	RoundsDir	string
	TanksDir	string
	ThemeDir	string
	AttemptInterval	time.Duration
	mux		*http.ServeMux
}

func (ctx *Instance) Initialize() error {
	if _, err := os.Stat(ctx.ThemeDir); err != nil {
		return err
	}

	if _, err := os.Stat(ctx.StateDir); err != nil {
		return err
	}

	ctx.mux = http.NewServeMux()

	ctx.Base = strings.TrimRight(ctx.Base, "/")

	ctx.RoundsDir = path.Join(ctx.StateDir, "rounds")
	ctx.TanksDir = path.Join(ctx.StateDir, "tanks")

	if _, err := os.Stat(ctx.RoundsDir); err != nil {
		os.Mkdir(ctx.RoundsDir, 0755)
	}

	if _, err := os.Stat(ctx.TanksDir); err != nil {
		os.Mkdir(ctx.TanksDir, 0755)
	}

	ctx.BindHandlers()

	return nil
}


func pathCleanse(parts []string) string {
	clean := make([]string, len(parts))
	for i := range parts {
		part := parts[i]
		part = strings.TrimLeft(part, ".")
		if p := strings.LastIndex(part, "/"); p >= 0 {
			part = part[p+1:]
		}
		clean[i] = part
	}
	return path.Join(clean...)
}

func (ctx *Instance) ThemePath(parts ...string) string {
	tail := pathCleanse(parts)
	return path.Join(ctx.ThemeDir, tail)
}

func (ctx *Instance) saveTank(tank Tank, token string) error {
	matched, err := regexp.MatchString("^[0-9a-zA-Z]{1,8}$", token)

	if err != nil || matched == false{
		return errors.New("Unrecognized token")
	}

	tank_base_path := path.Join(ctx.TanksDir, token)

	if _, err := os.Stat(tank_base_path); err != nil {
		os.Mkdir(tank_base_path, 0755)
	}

	log.Printf("Writing tank out to %s", tank_base_path)

	write_tank_component := func(component string, data string) {
		if f, err := os.OpenFile(path.Join(tank_base_path, component), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644); err == nil {
			defer f.Close()
			fmt.Fprint(f, data)
		}
		
	}

	write_tank_component("name", tank.name)
	write_tank_component("author", tank.author)
	write_tank_component("color", tank.color)
	write_tank_component("program", tank.program)

	for i, sensor := range tank.sensors {
		is_turret := 0
		if sensor.is_turret {
			is_turret = 1
		}

		write_tank_component(fmt.Sprintf("sensor%d", i), fmt.Sprintf("%d %d %d %d", sensor.radius, sensor.angle, sensor.width, is_turret))
	}
	
	return nil
}
