package main

import (
	"gym-cli/internal/handler"
)

func main() {
	app := handler.NewMenu()

	app.Run()
}
