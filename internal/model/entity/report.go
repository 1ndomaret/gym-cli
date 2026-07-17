package entity

import "time"

type IncomeEntry struct {
	PaymentDate time.Time
	InvoiceID   int
	Amount      float64
}

type MemberJoin struct {
	JoinDate  time.Time
	FirstName string
	LastName  string
}

type MvpMember struct {
	FirstName string
	LastName  string
	TotalPaid float64
}
