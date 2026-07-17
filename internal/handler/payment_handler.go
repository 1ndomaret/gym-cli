package handler

import (
	"bufio"
	"fmt"
	"gym-cli/internal/domain"
	"strconv"
	"strings"
)

type paymentHandler struct {
	puc    domain.PaymentUsecase
	iuc    domain.InvoiceUsecase
	reader *bufio.Reader
}

func NewPaymentHandler(puc domain.PaymentUsecase, iuc domain.InvoiceUsecase, reader *bufio.Reader) domain.PaymentHandler {
	return &paymentHandler{
		puc:    puc,
		iuc:    iuc,
		reader: reader,
	}
}

func (h *paymentHandler) Create() {
	for {
		fmt.Println("\n\033[0;33m========== CREATE PAYMENT ==========\033[0m")

		invoices, err := h.iuc.UnpaidInvoices()
		if err != nil {
			fmt.Println("\n\033[0;31mFailed to load report:", err, "\033[0m")
			return
		}

		if len(invoices) == 0 {
			fmt.Println("\nNo unpaid invoices — everyone is paid up.")
		} else {
			fmt.Printf("\n%-4s %-8s %-22s %14s  %-12s %s\n", "No", "InvoiceId", "Member", "Amount", "Due Date", "Status")
			fmt.Println(strings.Repeat("-", 75))
			var total float64
			for i, inv := range invoices {
				fmt.Printf("\033[0;33m%-4d\033[0m #%-7d %-22s %14.2f  %-12s %s\n",
					i+1,
					inv.InvoiceID,
					inv.FirstName+" "+inv.LastName,
					inv.Amount,
					inv.DueDate.Format("2006-01-02"),
					inv.InvoiceStatus)
				total += inv.Amount
			}
			fmt.Println("\n\033[0;33m0\033[0m    Back to menu ")
			fmt.Println(strings.Repeat("-", 75))
		}
		fmt.Print("\nInput Invoice number: ")

		numberInput, err := h.reader.ReadString('\n')
		if err != nil {
			fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
			continue
		}
		invoiceNumber, err := strconv.Atoi(strings.TrimSpace(numberInput))
		if err != nil {
			fmt.Println("\n\033[0;31mInvalid input, try again.\n\033[0m")
			continue
		}
		if invoiceNumber == 0 {
			return
		}

		invoiceNumber--
		if invoiceNumber < 0 || invoiceNumber > len(invoices)-1 {
			fmt.Println("\n\033[0;31mInvalid input, try again.\n\033[0m")
			continue
		}
		selectedInvoice := invoices[invoiceNumber]

		fmt.Printf("\nProcessing Payment for: \033[0;33m%s %s\033[0m\n", selectedInvoice.FirstName, selectedInvoice.LastName)
		fmt.Printf("Total Due: \033[0;32m%.2f\033[0m\n", selectedInvoice.Amount)
		fmt.Println(strings.Repeat("-", 55))

		methods, err := h.puc.PaymentMethods()
		methodSelection := 0
		for {
			fmt.Println("Select Payment Method:")
			for idx, method := range methods {
				fmt.Printf("\033[0;33m%d\033[0m. %s\n", idx+1, method.MethodName)
			}
			fmt.Printf("%-10s: ", "Choose Method")

			methodInput, err := h.reader.ReadString('\n')
			if err != nil {
				fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
				continue
			}

			methodSelection, err = strconv.Atoi(strings.TrimSpace(methodInput))
			methodSelection--
			if err != nil || methodSelection < 0 || methodSelection >= len(methods) {
				fmt.Println("\n\033[0;31mInvalid input, try again.\033[0m")
				continue
			}
			break
		}

		confirmPayment := false
		chosenMethod := methods[methodSelection]
		for {
			fmt.Printf("\nConfirm processing \033[0;32m%.2f\033[0m via \033[0;33m%s\033[0m? (yes/no): ", selectedInvoice.Amount, chosenMethod.MethodName)
			confirm, err := h.reader.ReadString('\n')
			if err != nil {
				fmt.Println("\n\033[0;31mUnexpected error occured when reading input.\033[0m", err)
				continue
			}
			confirm = strings.ToLower(strings.TrimSpace(confirm))

			if strings.ToLower(confirm) == "yes" {
				confirmPayment = true
			} else if strings.ToLower(confirm) != "no" {
				fmt.Println("\n\033[0;31mInvalid input, try again.\033[0m")
				continue
			}
			break
		}

		if !confirmPayment {
			fmt.Println("\n\033[0;31mPayment Cancelled.\033[0m")
			continue
		}

		err = h.puc.CreatePayment(&selectedInvoice, chosenMethod.PaymentMethodID)
		if err != nil {
			fmt.Println("\033[0;32mPayment successfully processed!\033[0m")
			continue
		}
		fmt.Println(strings.Repeat("-", 55))
		fmt.Println("Payment successfully processed!")
		fmt.Printf("Invoice #%d has been updated to PAID.\033[0m\n", selectedInvoice.InvoiceID)
		fmt.Println(strings.Repeat("-", 55))
		break
	}
}
