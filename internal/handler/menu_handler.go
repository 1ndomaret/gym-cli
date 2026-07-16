package handler

import (
	"fmt"
)

type Menu struct {
	// userHandler domain.UserHandler
}

func NewMenu() *Menu {
	return &Menu{
		// userHandler: uh,
	}
}

func (a *Menu) Run() {
	fmt.Println("Hello, World!")
}
