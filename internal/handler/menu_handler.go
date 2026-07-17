package handler

import (
	"bufio"
	"fmt"
	"gym-cli/internal/domain"
	"gym-cli/internal/model/entity"
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
		fmt.Print("========================================\n\033[0m")
		fmt.Printf(`
1. Login
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
			// a.login()
			a.adminMenu(&entity.User{})
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

func (a *menu) login() {
	for {
		user, err := a.userHandler.Login()
		if err != nil {
			fmt.Println("Here3")
			fmt.Println("\n\033[0;31m", err, "\n\033[0m")
			continue
		}

		switch user.UserType {
		case "admin":
			a.adminMenu(user)
		case "member":
			a.memberMenu(user)
		default:
			fmt.Println("\n\033[0;31m Unknown user type.\033[0m")
			continue
		}
		break
	}
}

func (a *menu) adminMenu(user *entity.User) {
	fmt.Printf("\n\033[0;32mLogged in as %s.\n\033[0m", user.Email)
	for {
		fmt.Println("\n\033[0;33m============= ADMIN MENU =============\033[0m")
		fmt.Printf(`
(MEMBER MANAGEMENT)
1. Add New Member
2. View All Members
3. Update Members
4. Delete Member

(REPORT)
5. Monthly Income Report
6. MVP Members (Top 10)
7. Member Joins

0. Logout

Choose Menu: `)

		menuInput, err := a.reader.ReadString('\n')
		if err != nil {
			fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
			return
		}

		menuSelection, err := strconv.Atoi(strings.TrimSpace(menuInput))

		if err != nil {
			fmt.Println("\n\033[0;31mInvalid input, try again.\033[0m")
			continue
		}

		exitApp := false
		switch menuSelection {
		case 1:
			a.userHandler.Create()
		case 2:
			a.userHandler.List()
		case 3:
			a.userHandler.Update()
		case 4:
		case 5:
			a.reportHandler.MonthlyIncome()
		case 6:
			a.reportHandler.MvpMembers()
		case 7:
			a.reportHandler.MemberJoins()
		case 0:
			exitApp = true
		default:
			fmt.Println("\n\033[0;31mInvalid input, try again.\033[0m")
			continue
		}

		if exitApp {
			fmt.Print("\033\n[0;32mLogged out.\n\n")
			break
		}
	}
}

func (a *menu) memberMenu(user *entity.User) {
	fmt.Printf("\n\033[0;32mLogged in as %s.\n\033[0m", user.Email)
	for {
		fmt.Println("\n\033[0;33m============= MEMBER MENU =============\033[0m")
		fmt.Printf(`0. Logout

Choose Menu: `)
		menuInput, err := a.reader.ReadString('\n')
		if err != nil {
			fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
			return
		}
		menuSelection, err := strconv.Atoi(strings.TrimSpace(menuInput))

		if err != nil {
			fmt.Println("\n\033[0;31mInvalid input, try again.\033[0m")
			continue
		}
		exitApp := false
		switch menuSelection {
		case 0:
			exitApp = true
		default:
			fmt.Println("\n\033[0;31mInvalid input, try again.\033[0m")
			continue
		}

		if exitApp {
			fmt.Print("\033\n[0;32mLogged out.\n\n")
			break
		}
	}
}
