package handler

import (
	"bufio"
	"gym-cli/internal/domain"
)

type invoiceHandler struct {
	uc     domain.InvoiceHandler
	reader *bufio.Reader
}

func NewInvoiceHandler(uc domain.InvoiceUsecase, reader *bufio.Reader) domain.InvoiceHandler {
	return &invoiceHandler{uc: uc, reader: reader}
}
