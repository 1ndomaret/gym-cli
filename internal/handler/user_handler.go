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

func NewUserHandler(uc domain.UserUsecase, reader *bufio.Reader) domain.UserHandler {
	return &userHandler{
		uc:     uc,
		reader: reader,
	}
}

func (h *userHandler) Create() {
	for {
		fmt.Println("\n\033[0;33m========== ADD NEW MEMBER ==========\033[0m")

		// ======== EMAIL ========
		fmt.Printf("%-15s: ", "Email")
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

		// ======== PASSWORD ========
		fmt.Printf("%-15s: ", "Password")
		passwordInput, err := h.reader.ReadString('\n')
		password := strings.TrimSpace(passwordInput)
		if err != nil {
			fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
			continue
		}

		// ======== FIRSTNAME ========
		fmt.Printf("%-15s: ", "First Name")
		firstNameInput, err := h.reader.ReadString('\n')
		firstName := strings.TrimSpace(firstNameInput)
		if err != nil {
			fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
			continue
		}

		// ======== LASTNAME ========
		fmt.Printf("%-15s: ", "Last Name")
		lastNameInput, err := h.reader.ReadString('\n')
		lastName := strings.TrimSpace(lastNameInput)
		if err != nil {
			fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
			continue
		}

		// ======== MEMBER TIER ========
		var tier string
		var memberTierId int
		for {
			// TODO: PRINT TIERS DYNAMICALLY
			fmt.Printf(`Tiers:
  1. Gold
  2. Silver
  3. Bronze	
%-15s: `, "Select Tier")
			tierInput, err := h.reader.ReadString('\n')
			if err != nil {
				fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
				continue
			}
			memberTierId, err = strconv.Atoi(strings.TrimSpace(tierInput))
			if err != nil {
				fmt.Println("\n\033[0;31mInvalid input, try again.\n\033[0m")
				fmt.Printf("Email: %s\n", email)
				continue
			}

			switch memberTierId {
			case 1:
				tier = "Gold"
			case 2:
				tier = "Silver"
			case 3:
				tier = "Bronze"
			default:
				fmt.Println("\n\033[0;31mInvalid input, try again.\n\033[0m")
				continue
			}
			break
		}

		// ======== ADDRESS ========
		fmt.Printf("%-15s: ", "Address")
		addressInput, err := h.reader.ReadString('\n')
		address := strings.TrimSpace(addressInput)
		if err != nil {
			fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
			continue
		}

		// TODO: ADD INSERT USER

		err = h.uc.Create(email, password, firstName, lastName, address, memberTierId)
		if err != nil {
			fmt.Println("\n\033[0;31m", err, "\033[0m")
			return
		}

		fmt.Println("\n\033[0;32mNew Member added.\033[0m")
		fmt.Printf("%-15s: %s %s\n", "Name", firstName, lastName)
		fmt.Printf("%-15s: %s\n", "Email", email)
		fmt.Printf("%-15s: %s\n", "Tier", tier)

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
		fmt.Print("\n")
		break
	}
}

func (h *userHandler) List() {
	fmt.Println("\n\033[0;33m============= MEMBER LIST =============\033[0m")
	memberList, err := h.uc.MemberList()
	if err != nil {
		fmt.Println("\n\033[0;31m", err, "\033[0m")
		return
	}

	fmt.Printf("%-3s %-20s %-30s %-10s %-15s %-10s\n",
		"No", "Name", "Email", "Tier", "Joined", "Status")
	fmt.Println(strings.Repeat("-", 90))

	for i, member := range memberList {
		memberName := member.FirstName + " " + member.LastName
		status := "Inactive"
		if member.Status {
			status = "Active"
		}
		fmt.Printf("%-3d %-20s %-30s %-10s %-15s %-10s\n",
			i+1,
			memberName,
			member.Email,
			member.TierName,
			member.CreatedAt.Format("2006-01-02"),
			status,
		)
	}
	fmt.Print("\n\033[0;32mPress (Enter) to continue.\033[0m")
	_, err = h.reader.ReadString('\n')
	if err != nil {
		fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
		return
	}
	fmt.Print("\n")
}
