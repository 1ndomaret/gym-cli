package entity

import "time"

type UserDetail struct {
	Email     string
	FirstName string
	LastName  string
	TierName  string
	Address   string
	CreatedAt time.Time
	Status    bool
}
