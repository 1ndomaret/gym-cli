package handler

import (
	"bufio"
	"fmt"
	"gym-cli/internal/domain"
	"gym-cli/internal/model/entity"
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
		fmt.Println("\n\033[0;33m========== ADD NEW MEMBER ==========\n\033[0m")

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
			fmt.Println("\n\033[0;31mInvalid email format, please try again.\n\033[0m")
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
  1. Bronze
  2. Silver
  3. Gold	
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
				tier = "Bronze"
			case 2:
				tier = "Silver"
			case 3:
				tier = "Gold"
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
				fmt.Println("\n\033[0;31mInvalid input, try again.\n\033[0m")
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
	fmt.Println("\n\033[0;33m============= MEMBER LIST =============\n\033[0m")
	memberList, err := h.uc.MemberList()
	if err != nil {
		fmt.Printf("\n\033[0;31m%s\033[0m", err)
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

func (h *userHandler) Login() (*entity.User, error) {
	var email string
	var password string

	for {
		fmt.Print("\n\033[0;33mPlease Enter your Credentials\n\033[0m")
		fmt.Printf("%-10s: ", "Email")
		emailInput, err := h.reader.ReadString('\n')
		email = strings.TrimSpace(emailInput)
		if err != nil {
			fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
			continue
		}

		fmt.Printf("%-10s: ", "Password")
		passwordInput, err := h.reader.ReadString('\n')
		password = strings.TrimSpace(passwordInput)
		if err != nil {
			fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
			continue
		}
		break
	}
	user, err := h.uc.Login(email, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (h *userHandler) Update() {
	for {
		fmt.Println("\n\033[0;33m============= UPDATE MEMBER =============\033\n[0m")
		memberList, err := h.uc.MemberList()
		if err != nil {
			fmt.Printf("\n\033[0;31m%s\033[0m", err)
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
		fmt.Println("\n0. Back to menu ")

		fmt.Print("\nInput member number: ")

		numberInput, err := h.reader.ReadString('\n')
		if err != nil {
			fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
			continue
		}
		memberNumber, err := strconv.Atoi(strings.TrimSpace(numberInput))
		if err != nil {
			fmt.Println("\n\033[0;31mInvalid input, try again.\n\033[0m")
			continue
		}
		if memberNumber == 0 {
			return
		}

		memberNumber--
		if memberNumber < 0 || memberNumber > len(memberList)-1 {
			fmt.Println("\n\033[0;31mInvalid input, try again.\n\033[0m")
			continue
		}

		selectedMember := memberList[memberNumber]
		fmt.Printf("\nUpdating Details for: \033[0;33m%s %s\033[0m\n", selectedMember.FirstName, selectedMember.LastName)
		fmt.Println("\033[0;32mLeave empty and press (Enter) to keep the current value.\033[0m")
		fmt.Println(strings.Repeat("-", 50))

		// ======== FIRST NAME ========
		fmt.Printf("First Name [%s]: ", selectedMember.FirstName)
		newFirstName, _ := h.reader.ReadString('\n')
		newFirstName = strings.TrimSpace(newFirstName)
		if newFirstName == "" {
			newFirstName = selectedMember.FirstName
		}

		// ======== LAST NAME ========
		fmt.Printf("Last Name [%s]: ", selectedMember.LastName)
		newLastName, _ := h.reader.ReadString('\n')
		newLastName = strings.TrimSpace(newLastName)
		if newLastName == "" {
			newLastName = selectedMember.LastName
		}

		// ======== EMAIL ========
		var newEmail string
		for {
			fmt.Printf("Email [%s]: ", selectedMember.Email)
			newEmail, _ = h.reader.ReadString('\n')
			newEmail = strings.TrimSpace(newEmail)
			if newEmail == "" {
				newEmail = selectedMember.Email
			}

			addr, err := mail.ParseAddress(newEmail)
			if err != nil || addr.Address != newEmail {
				fmt.Println("\n\033[0;31mInvalid email format, please try again.\n\033[0m")
				continue
			} else if newEmail == "test@mail.com" {
				// TODO: ADD DUPLICATE CHECKER
				fmt.Println("\n\033[0;31mEmail is taken, please enter another email.\n\033[0m")
				continue
			}

			break
		}

		// ======== ADDRESS ========
		fmt.Printf("Address [%s]: ", selectedMember.Address)
		newAddress, _ := h.reader.ReadString('\n')
		newAddress = strings.TrimSpace(newAddress)
		if newAddress == "" {
			newAddress = selectedMember.Address
		}

		// ======== STATUS ========
		currentStatusStr := "Active"
		if !selectedMember.Status {
			currentStatusStr = "Inactive"
		}
		newStatus := selectedMember.Status

		for {
			fmt.Printf("Status (active/inactive) [%s]: ", currentStatusStr)
			newStatusInput, _ := h.reader.ReadString('\n')
			newStatusInput = strings.ToLower(strings.TrimSpace(newStatusInput))

			switch newStatusInput {
			case "active":
				newStatus = true
			case "inactive":
				newStatus = false
			case "":
			default:
				fmt.Println("\n\033[0;31mInvalid input, please try again.\n\033[0m")
				continue
			}
			break
		}

		selectedMember.FirstName = newFirstName
		selectedMember.LastName = newLastName
		selectedMember.Email = newEmail
		selectedMember.Status = newStatus
		selectedMember.Address = newAddress

		err = h.uc.UpdateMember(selectedMember.UserId, &selectedMember)
		if err != nil {
			fmt.Printf("\n\033[0;31m%s\033[0m", err)
			return
		}

		fmt.Println("\n\033[0;32mMember Successfully Updated.\033[0m")

		anotherMember := false
		for {
			fmt.Print("\nWould you like to update another member? (yes/no) ")
			cont, err := h.reader.ReadString('\n')
			cont = strings.ToLower(strings.TrimSpace(cont))
			if err != nil {
				fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
				continue
			}

			if strings.ToLower(cont) == "yes" {
				anotherMember = true
			} else if strings.ToLower(cont) != "no" {
				fmt.Println("\n\033[0;31mInvalid input, try again.\n\033[0m")
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

// User View
func (h *userHandler) ViewSchedule(userID int) {
	fmt.Println("\n\033[0;33m========== UPCOMING SCHEDULE ==========\033[0m")

	events, err := h.uc.ViewSchedule(userID)
	if err != nil {
		fmt.Println("\n\033[0;31mFailed to load schedule:", err, "\033[0m")
		return
	}

	if len(events) == 0 {
		fmt.Println("\nNo upcoming events available for your tier.")
	} else {
		fmt.Printf("%-20s %s\n", "When", "Event")
		fmt.Println(strings.Repeat("-", 50))
		for _, e := range events {
			fmt.Printf("%-20s %s\n", e.Schedule.Format("2006-01-02 15:04"), e.EventName)
		}
	}

	fmt.Print("\n\033[0;32mPress (Enter) to continue.\033[0m")
	h.reader.ReadString('\n')
}

func (h *userHandler) ViewPendingPayment(userID int) {
	fmt.Println("\n\033[0;33m========== PENDING PAYMENTS ==========\033[0m")

	invoices, err := h.uc.ViewPendingPayment(userID)
	if err != nil {
		fmt.Println("\n\033[0;31mFailed to load payments:", err, "\033[0m")
		return
	}

	if len(invoices) == 0 {
		fmt.Println("\nYou have no pending payments.")
	} else {
		fmt.Printf("%-10s %15s  %-12s %s\n", "Invoice", "Amount", "Due Date", "Status")
		fmt.Println(strings.Repeat("-", 52))
		var total float64
		for _, inv := range invoices {
			fmt.Printf("#%-9d %15.2f  %-12s %s\n",
				inv.InvoiceID, inv.Amount, inv.DueDate.Format("2006-01-02"), inv.InvoiceStatus)
			total += inv.Amount
		}
		fmt.Println(strings.Repeat("-", 52))
		fmt.Printf("%-10s %15.2f\n", "TOTAL", total)
	}

	fmt.Print("\n\033[0;32mPress (Enter) to continue.\033[0m")
	h.reader.ReadString('\n')
}
