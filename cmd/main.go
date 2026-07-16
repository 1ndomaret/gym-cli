package main

import (
	"bufio"
	"gym-cli/internal/handler"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	uh := handler.NewUserHandler(reader)
	app := handler.NewMenu(uh, reader)

	app.Run()
}
