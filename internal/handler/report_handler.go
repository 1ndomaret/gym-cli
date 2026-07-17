package handler

import (
	"bufio"
	"fmt"
	"gym-cli/internal/domain"
	"strings"
)

type reportHandler struct {
	uc     domain.ReportUsecase
	reader *bufio.Reader
}

func NewReportHandler(uc domain.ReportUsecase, reader *bufio.Reader) domain.ReportHandler {
	return &reportHandler{uc: uc, reader: reader}
}

func (h *reportHandler) MonthlyIncome() {
	fmt.Println("\n\033[0;33m========== MONTHLY INCOME REPORT ==========\033[0m")

	report, err := h.uc.MonthlyIncome()
	if err != nil {
		fmt.Println("\n\033[0;31mFailed to load report:", err, "\033[0m")
		return
	}

	if len(report) == 0 {
		fmt.Println("\nNo income recorded yet.")
		fmt.Print("\nPress (Enter) to continue.")
		h.reader.ReadString('\n')
		return
	}

	var grandTotal, monthTotal float64
	currentMonth := ""

	for _, row := range report {
		month := row.PaymentDate.Format("January 2006")

		if month != currentMonth {
			if currentMonth != "" {
				fmt.Printf("  %-21s %12.2f\n", "Total for "+currentMonth+":", monthTotal)
			}
			fmt.Printf("\n%s\n", month)
			currentMonth = month
			monthTotal = 0
		}

		fmt.Printf("  %-12s %-8s %12.2f\n",
			row.PaymentDate.Format("2006-01-02"),
			fmt.Sprintf("#%d", row.InvoiceID),
			row.Amount)

		monthTotal += row.Amount
		grandTotal += row.Amount
	}

	fmt.Printf("  %-21s %12.2f\n", "Total for "+currentMonth+":", monthTotal)
	fmt.Printf("\n%-23s %12.2f\n", "GRAND TOTAL:", grandTotal)

	fmt.Print("\nPress (Enter) to continue.")
	h.reader.ReadString('\n')
}

func (h *reportHandler) MvpMembers() {
	fmt.Println("\n\033[0;33m========== TOP 10 MVP MEMBERS ==========\033[0m")

	members, err := h.uc.MvpMembers()
	if err != nil {
		fmt.Println("\n\033[0;31mFailed to load report:", err, "\033[0m")
		return
	}

	if len(members) == 0 {
		fmt.Println("\nNo paying members yet.")
	} else {
		fmt.Printf("%-5s %-25s %15s\n", "Rank", "Member", "Total Paid")
		fmt.Println(strings.Repeat("-", 47))
		for i, m := range members {
			fmt.Printf("%-5d %-25s %15.2f\n", i+1, m.FirstName+" "+m.LastName, m.TotalPaid)
		}
	}

	fmt.Print("\nPress (Enter) to continue.")
	h.reader.ReadString('\n')
}

func (h *reportHandler) MemberJoins() {
	fmt.Println("\n\033[0;33m========== MEMBER JOINS BY MONTH ==========\033[0m")

	joins, err := h.uc.MemberJoins()
	if err != nil {
		fmt.Println("\n\033[0;31mFailed to load report:", err, "\033[0m")
		return
	}

	if len(joins) == 0 {
		fmt.Println("\nNo members yet.")
		fmt.Print("\nPress (Enter) to continue.")
		h.reader.ReadString('\n')
		return
	}

	grandTotal := 0
	monthCount := 0
	currentMonth := ""

	for _, j := range joins {
		month := j.JoinDate.Format("January 2006")

		if month != currentMonth {
			if currentMonth != "" {
				fmt.Printf("  %-25s %6d\n", "Total for "+currentMonth+":", monthCount)
			}
			fmt.Printf("\n%s\n", month)
			currentMonth = month
			monthCount = 0
		}

		fmt.Printf("  %-12s %s\n", j.JoinDate.Format("2006-01-02"), j.FirstName+" "+j.LastName)
		monthCount++
		grandTotal++
	}

	fmt.Printf("  %-25s %6d\n", "Total for "+currentMonth+":", monthCount)
	fmt.Printf("\n%-27s %6d\n", "GRAND TOTAL:", grandTotal)

	fmt.Print("\nPress (Enter) to continue.")
	h.reader.ReadString('\n')
}
