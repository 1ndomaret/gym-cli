package db

import (
	"database/sql"
	"gym-cli/internal/domain"
)

type invoiceRepository struct {
	db *sql.DB
}

func NewInvoiceRepository(db *sql.DB) domain.InvoiceRepository {
	return &invoiceRepository{
		db: db,
	}
}
