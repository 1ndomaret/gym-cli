package entity

import "time"

type Payment struct {
	PaymentID       int
	InvoiceID       int
	PaymentDate     time.Time
	PaymentMethodID int
	PaymentStatus   string
}
