package db

import (
	"context"
	"database/sql"
	"gym-cli/internal/domain"
	"gym-cli/internal/model/entity"
)

type paymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) domain.PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

func (r *paymentRepository) CreatePayment(ctx context.Context, invoice *entity.UnpaidInvoice, methodId int) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// ===== PAYMENT QUERY =====
	paymentQuery := `
		INSERT INTO Payments (InvoiceId, PaymentMethodId, PaymentDate, PaymentStatus) 
			VALUES (?, ?, NOW(), 'completed')
	`
	_, err = tx.ExecContext(ctx, paymentQuery, invoice.InvoiceID, methodId)
	if err != nil {
		return err
	}

	// ===== INVOICE QUERY =====
	invoiceQuery := `
		UPDATE Invoices 
		SET InvoiceStatus = 'paid' 
		WHERE InvoiceId = ?
	`
	_, err = tx.ExecContext(ctx, invoiceQuery, invoice.InvoiceID)
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (r *paymentRepository) PaymentMethods(ctx context.Context) ([]entity.PaymentMethod, error) {
	query := `
		SELECT PaymentMethodId, MethodName 
			FROM PaymentMethod
			ORDER BY PaymentMethodId ASC
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var methods []entity.PaymentMethod
	for rows.Next() {
		var method entity.PaymentMethod
		err := rows.Scan(
			&method.PaymentMethodID,
			&method.MethodName,
		)
		if err != nil {
			return nil, err
		}
		methods = append(methods, method)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return methods, nil
}
