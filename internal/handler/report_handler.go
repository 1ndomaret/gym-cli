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
	} else {
		fmt.Printf("%-10s %18s\n", "Month", "Total Income")
		fmt.Println(strings.Repeat("-", 29))
		var grandTotal float64
		for _, row := range report {
			fmt.Printf("%-10s %18.2f\n", row.Month, row.TotalIncome)
			grandTotal += row.TotalIncome
		}
		fmt.Println(strings.Repeat("-", 29))
		fmt.Printf("%-10s %18.2f\n", "TOTAL", grandTotal)
	}

	fmt.Print("\nPress (Enter) to continue.")
	h.reader.ReadString('\n')
}
