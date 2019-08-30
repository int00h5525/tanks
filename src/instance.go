package main

import (
	//"fmt"
	"strings"
	"time"
	"log"
	"net/http"
	//"os"
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
	log.Print("Initializing context")

	ctx.mux = http.NewServeMux()

	ctx.Base = strings.TrimRight(ctx.Base, "/")

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
