package usecase

import (
	"context"
	"gym-cli/internal/domain"
	"gym-cli/internal/model/entity"
)

type paymentUsecase struct {
	repo domain.PaymentRepository
}

func NewPaymentUsecase(repo domain.PaymentRepository) domain.PaymentUsecase {
	return &paymentUsecase{
		repo: repo,
	}
}

func (u *paymentUsecase) CreatePayment(invoice *entity.UnpaidInvoice, methodId int) error {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, setTimeout)
	defer cancel()
	return u.repo.CreatePayment(ctx, invoice, methodId)
}

func (u *paymentUsecase) PaymentMethods() ([]entity.PaymentMethod, error) {
	ctx := context.TODO()
	ctx, cancel := context.WithTimeout(ctx, setTimeout)
	defer cancel()
	return u.repo.PaymentMethods(ctx)
}
