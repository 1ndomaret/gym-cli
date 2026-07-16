package handler

import (
	"bufio"
	"fmt"
	"gym-cli/internal/domain"
	"strconv"
	"strings"
)

type menu struct {
	userHandler   domain.UserHandler
	reportHandler domain.ReportHandler
	reader        *bufio.Reader
}

func NewMenu(uh domain.UserHandler, rh domain.ReportHandler, reader *bufio.Reader) *menu {
	return &menu{
		userHandler:   uh,
		reportHandler: rh,
		reader:        reader,
	}
}

func (a *menu) Run() {

	for {
		fmt.Println("\033[0;33m========================================")
		fmt.Println("     GOLD'S GYM MEMBERSHIP SYSTEM")
		fmt.Print("========================================\033[0m")
		fmt.Printf(`
1. Add New Member
2. View All Members
3. Update Members
4. Delete Member
5. Monthly Income Report

0. Exit Application

Choose Menu: `)

		menuInput, err := a.reader.ReadString('\n')
		if err != nil {
			fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
			return
		}

		menuSelection, err := strconv.Atoi(strings.TrimSpace(menuInput))

		if err != nil {
			fmt.Println("\n\033[0;31mInvalid input, try again.\n\033[0m")
			continue
		}

		exitApp := false
		switch menuSelection {
		case 1:
			a.userHandler.Create()
		case 2:
			a.userHandler.List()
		case 5:
			a.reportHandler.MonthlyIncome()
		case 0:
			exitApp = true
		default:
			fmt.Println("\n\033[0;31mInvalid input, try again.\n\033[0m")
			continue
		}

		if exitApp {
			fmt.Println("\033[0;33mExiting Application.")
			break
		}
	}
}
