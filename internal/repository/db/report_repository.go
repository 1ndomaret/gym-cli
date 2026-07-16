package db

import (
	"context"
	"database/sql"
	"gym-cli/internal/domain"
	"gym-cli/internal/model/entity"
)

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) domain.ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) MonthlyIncome(ctx context.Context) ([]entity.MonthlyIncome, error) {
	query := `
		SELECT DATE_FORMAT(p.PaymentDate, '%Y-%m') AS Month,
		       COALESCE(SUM(i.Amount), 0)          AS TotalIncome
		FROM Payments p
		JOIN Invoices i ON i.InvoiceId = p.InvoiceId
		WHERE p.PaymentStatus = 'completed'
		GROUP BY DATE_FORMAT(p.PaymentDate, '%Y-%m')
		ORDER BY Month
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var report []entity.MonthlyIncome
	for rows.Next() {
		var row entity.MonthlyIncome
		if err := rows.Scan(&row.Month, &row.TotalIncome); err != nil {
			return nil, err
		}
		report = append(report, row)
	}
	return report, rows.Err()
}
