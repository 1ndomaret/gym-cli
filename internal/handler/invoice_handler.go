package handler

import (
	"bufio"
	"fmt"
	"gym-cli/internal/domain"
	"strings"
)

type invoiceHandler struct {
	uc     domain.InvoiceUsecase
	reader *bufio.Reader
}

func NewInvoiceHandler(uc domain.InvoiceUsecase, reader *bufio.Reader) domain.InvoiceHandler {
	return &invoiceHandler{uc: uc, reader: reader}
}

func (h *invoiceHandler) UnpaidInvoices() {
	fmt.Println("\n\033[0;33m========== CREATE PAYMENT ==========\033[0m")

	invoices, err := h.uc.UnpaidInvoices()
	if err != nil {
		fmt.Println("\n\033[0;31mFailed to load report:", err, "\033[0m")
		return
	}

	if len(invoices) == 0 {
		fmt.Println("\nNo unpaid invoices — everyone is paid up.")
	} else {
		fmt.Printf("\n%-8s %-22s %14s  %-12s %s\n", "InvoiceId", "Member", "Amount", "Due Date", "Status")
		fmt.Println(strings.Repeat("-", 70))
		var total float64
		for _, inv := range invoices {
			fmt.Printf("#%-7d %-22s %14.2f  %-12s %s\n",
				inv.InvoiceID,
				inv.FirstName+" "+inv.LastName,
				inv.Amount,
				inv.DueDate.Format("2006-01-02"),
				inv.InvoiceStatus)
			total += inv.Amount
		}
		fmt.Println(strings.Repeat("-", 70))
		fmt.Printf("%-31s %14.2f\n", "TOTAL VALUE:", total)
	}

	fmt.Print("\n\033[0;32mPress (Enter) to continue.\033[0m")
	h.reader.ReadString('\n')
}
