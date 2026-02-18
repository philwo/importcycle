package main

import (
	"github.com/philwo/importcycle/internal/app"
)

func main() {
	a := &app.App{Name: "ImportCycleDemo"}
	a.Run()
}
