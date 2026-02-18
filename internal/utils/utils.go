package utils

import (
	"fmt"
	"github.com/philwo/importcycle/internal/app"
)

func PrintHello() {
	fmt.Println("Hello from utils!")
	// This is where the cycle happens: utility depends on app
	fmt.Println("Utility knows app name is:", app.GetAppName())
}
