package usecase

import (
	"gym-cli/internal/domain"
)

type paymentUsecase struct {
	repo domain.PaymentRepository
}

func NewPaymentUsecase(repo domain.PaymentRepository) domain.PaymentUsecase {
	return &paymentUsecase{
		repo: repo,
	}
}
