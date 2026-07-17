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

func (r *reportRepository) MonthlyIncome(ctx context.Context) ([]entity.IncomeEntry, error) {
	query := `
		SELECT p.PaymentDate, p.InvoiceId, i.Amount
		FROM Payments p
		JOIN Invoices i ON i.InvoiceId = p.InvoiceId
		WHERE p.PaymentStatus = 'completed'
		ORDER BY p.PaymentDate
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var report []entity.IncomeEntry
	for rows.Next() {
		var row entity.IncomeEntry
		if err := rows.Scan(&row.PaymentDate, &row.InvoiceID, &row.Amount); err != nil {
			return nil, err
		}
		report = append(report, row)
	}
	return report, rows.Err()
}

func (r *reportRepository) MvpMembers(ctx context.Context) ([]entity.MvpMember, error) {
	query := `
		SELECT up.FirstName, up.LastName, SUM(i.Amount) AS TotalPaid
		FROM Payments p
		JOIN Invoices i      ON i.InvoiceId = p.InvoiceId
		JOIN UserProfiles up ON up.UserProfileId = i.UserProfileId
		WHERE p.PaymentStatus = 'completed'
		GROUP BY up.UserProfileId, up.FirstName, up.LastName
		ORDER BY TotalPaid DESC
		LIMIT 10
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []entity.MvpMember
	for rows.Next() {
		var m entity.MvpMember
		if err := rows.Scan(&m.FirstName, &m.LastName, &m.TotalPaid); err != nil {
			return nil, err
		}
		members = append(members, m)
	}
	return members, rows.Err()
}

func (r *reportRepository) MemberJoins(ctx context.Context) ([]entity.MemberJoin, error) {
	query := `
		SELECT CreatedAt, FirstName, LastName
		FROM UserProfiles
		ORDER BY CreatedAt
			`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var joins []entity.MemberJoin
	for rows.Next() {
		var j entity.MemberJoin
		if err := rows.Scan(&j.JoinDate, &j.FirstName, &j.LastName); err != nil {
			return nil, err
		}
		joins = append(joins, j)
	}
	return joins, rows.Err()
}
