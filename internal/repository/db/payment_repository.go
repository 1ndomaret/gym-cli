package db

import (
	"database/sql"
	"gym-cli/internal/domain"
)

type paymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) domain.PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}
