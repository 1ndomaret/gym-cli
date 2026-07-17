package usecase

import (
	"gym-cli/internal/domain"
)

type invoiceUsecase struct {
	repo domain.InvoiceRepository
}

func NewInvoiceUsecase(repo domain.InvoiceRepository) domain.InvoiceUsecase {
	return &invoiceUsecase{
		repo: repo,
	}
}
