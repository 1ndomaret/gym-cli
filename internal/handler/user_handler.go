package handler

import (
	"bufio"
	"fmt"
	"gym-cli/internal/domain"
	"net/mail"
	"strconv"
	"strings"
)

type userHandler struct {
	uc     domain.UserUsecase
	reader *bufio.Reader
}

func NewUserHandler(reader *bufio.Reader) domain.UserHandler {
	return &userHandler{
		reader: reader,
	}
}

func (h *userHandler) Create() {
	for {
		fmt.Println("\n\033[0;33m========== ADD NEW MEMBER ==========\033[0m")
		fmt.Print("Email: ")
		emailInput, err := h.reader.ReadString('\n')
		if err != nil {
			fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
			continue
		}

		email := strings.TrimSpace(emailInput)
		addr, err := mail.ParseAddress(email)
		if err != nil || addr.Address != email {
			fmt.Println("\n\033[0;31mInvalid email format, please try again.\033[0m")
			continue
		} else if email == "test@mail.com" {
			// TODO: ADD DUPLICATE CHECKER
			fmt.Println("\n\033[0;31mEmail is taken, please enter another email.\033[0m")
			continue
		}

		var tier string
		for {
			// TODO: PRINT TIERS DYNAMICALLY
			fmt.Printf(`
Tiers:
1. Gold
2. Silver
3. Bronze	
				
Select tier: `)
			tierInput, err := h.reader.ReadString('\n')
			if err != nil {
				fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
				continue
			}
			tierSelection, err := strconv.Atoi(strings.TrimSpace(tierInput))
			if err != nil {
				fmt.Println("\n\033[0;31mInvalid input, try again.\033[0m")
				continue
			}

			switch tierSelection {
			case 1:
				tier = "Gold"
			case 2:
				tier = "Silver"
			case 3:
				tier = "Bronze"
			default:
				fmt.Println("\n\033[0;31mInvalid input, try again.\033[0m")
				continue
			}
			break
		}

		// TODO: ADD INSERT USER
		fmt.Println("\n\033[0;32mNew Member added.\033[0m")
		fmt.Println("Email:", email)
		fmt.Println("Tier:", tier)

		anotherMember := false
		for {
			fmt.Print("\nWould you like to add another member? (yes/no) ")
			cont, err := h.reader.ReadString('\n')
			cont = strings.ToLower(strings.TrimSpace(cont))
			if err != nil {
				fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
				continue
			}

			if strings.ToLower(cont) == "yes" {
				anotherMember = true
			} else if strings.ToLower(cont) != "no" {
				fmt.Println("\n\033[0;31mInvalid input, try again.\033[0m")
				continue
			}
			break
		}
		if anotherMember {
			continue
		}
		break
	}
}

func (h *userHandler) List() {

}
