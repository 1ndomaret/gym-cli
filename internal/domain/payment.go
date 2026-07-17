package domain

import (
	"context"
	"gym-cli/internal/model/entity"
)

type PaymentRepository interface {
	CreatePayment(ctx context.Context, invoice *entity.UnpaidInvoice, methodId int) error
	PaymentMethods(ctx context.Context) ([]entity.PaymentMethod, error)
}

type PaymentUsecase interface {
	CreatePayment(invoice *entity.UnpaidInvoice, methodId int) error
	PaymentMethods() ([]entity.PaymentMethod, error)
}

type PaymentHandler interface {
	Create()
}
