package app

import (
	"fmt"
	"github.com/philwo/importcycle/internal/utils"
)

type App struct {
	Name string
}

func (a *App) Run() {
	fmt.Println("Running app:", a.Name)
	utils.PrintHello()
}

func GetAppName() string {
	return "Demonstration App"
}
