package usecase

import (
	"context"
	"gym-cli/internal/domain"
	"gym-cli/internal/model/entity"
	"time"
)

type invoiceUsecase struct {
	repo domain.InvoiceRepository
}

func NewInvoiceUsecase(repo domain.InvoiceRepository) domain.InvoiceUsecase {
	return &invoiceUsecase{
		repo: repo,
	}
}

func (u *invoiceUsecase) UnpaidInvoices() ([]entity.UnpaidInvoice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return u.repo.UnpaidInvoices(ctx)
}
