package domain

import (
	"context"
	"gym-cli/internal/model/entity"
)

type InvoiceRepository interface {
	UnpaidInvoices(ctx context.Context) ([]entity.UnpaidInvoice, error)
}

type InvoiceUsecase interface {
	UnpaidInvoices() ([]entity.UnpaidInvoice, error)
}

type InvoiceHandler interface {
	UnpaidInvoices()
}
