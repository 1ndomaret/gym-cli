package entity

import "time"

type Invoice struct {
	InvoiceID     int
	UserProfileID int
	MemberTierID  int
	Amount        float64
	DueDate       time.Time
	InvoiceStatus string
}
