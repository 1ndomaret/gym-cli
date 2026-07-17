package handler

import (
	"bufio"
	"gym-cli/internal/domain"
)

type paymentHandler struct {
	uc     domain.PaymentHandler
	reader *bufio.Reader
}

func NewPaymentHandler(uc domain.PaymentUsecase, reader *bufio.Reader) domain.PaymentHandler {
	return &paymentHandler{uc: uc, reader: reader}
}
