package handler

import (
	"bufio"
	"fmt"
	"gym-cli/internal/domain"
	"gym-cli/internal/model/entity"
	"net/mail"
	"strconv"
	"strings"
	"time"
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
			fmt.Printf(`Tiers:
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
				fmt.Println("\n\033[0;31mInvalid input, try again.\n\033[0m")
				fmt.Printf("Email: %s\n", email)
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
				fmt.Println("\n\033[0;31mInvalid input, try again.\n\033[0m")
				fmt.Printf("Email: %s\n", email)
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
		fmt.Print("\n")
		break
	}
}

func (h *userHandler) List() {
	fmt.Println("\n\033[0;33m============= MEMBER LIST =============\033[0m")
	users := []entity.UserDetails{
		{
			UserID:        1,
			Email:         "john.doe@example.com",
			UserType:      "Customer",
			UserProfileID: 101,
			MemberTierID:  1,
			FirstName:     "John",
			LastName:      "Doe",
			Address:       "123 Main Street",
			CreatedAt:     time.Date(2026, time.January, 15, 10, 30, 0, 0, time.UTC),
			Status:        "Active",
		},
		{
			UserID:        2,
			Email:         "jane.smith@example.com",
			UserType:      "Customer",
			UserProfileID: 102,
			MemberTierID:  2,
			FirstName:     "Jane",
			LastName:      "Smith",
			Address:       "456 Oak Avenue",
			CreatedAt:     time.Date(2026, time.March, 20, 14, 45, 0, 0, time.UTC),
			Status:        "Active",
		},
	}

	tiers := []entity.Tier{
		{
			TierId:      1,
			TierName:    "Gold",
			MonthlyCost: 30.00,
		},
		{
			TierId:      2,
			TierName:    "Silver",
			MonthlyCost: 20.00,
		},
		{
			TierId:      3,
			TierName:    "Bronze",
			MonthlyCost: 10.00,
		},
	}

	tierMap := make(map[int]entity.Tier)

	for _, tier := range tiers {
		tierMap[tier.TierId] = tier
	}

	fmt.Printf("%-3s %-20s %-30s %-10s %-10s\n",
		"No", "Name", "Email", "Tier", "Status")
	fmt.Println(strings.Repeat("-", 80))

	for i, user := range users {
		tier := tierMap[user.MemberTierID]

		name := user.FirstName + " " + user.LastName

		fmt.Printf("%-3d %-20s %-30s %-10s %-10s\n",
			i+1,
			name,
			user.Email,
			tier.TierName,
			user.Status,
		)
	}
	fmt.Print("\nPress (Enter) to continue.")
	_, err := h.reader.ReadString('\n')
	if err != nil {
		fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
		return
	}
	fmt.Print("\n")
}
