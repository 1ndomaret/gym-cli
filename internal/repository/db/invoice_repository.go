package db

import (
	"context"
	"database/sql"
	"gym-cli/internal/domain"
	"gym-cli/internal/model/entity"
)

type invoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) domain.InvoiceRepository {
	return &invoiceRepository{
		db: db,
	}
}

func (r *invoiceRepository) UnpaidInvoices(ctx context.Context) ([]entity.UnpaidInvoice, error) {
	query := `
		SELECT i.InvoiceId, up.FirstName, up.LastName, i.Amount, i.DueDate, i.InvoiceStatus
		FROM Invoices i
		JOIN UserProfiles up ON up.UserProfileId = i.UserProfileId
		WHERE i.InvoiceStatus IN ('pending', 'overdue')
		ORDER BY i.DueDate
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []entity.UnpaidInvoice
	for rows.Next() {
		var inv entity.UnpaidInvoice
		if err := rows.Scan(&inv.InvoiceID, &inv.FirstName, &inv.LastName,
			&inv.Amount, &inv.DueDate, &inv.InvoiceStatus); err != nil {
			return nil, err
		}
		invoices = append(invoices, inv)
	}
	return invoices, rows.Err()
}
