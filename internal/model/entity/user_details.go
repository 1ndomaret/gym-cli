package entity

import "time"

type UserDetails struct {
	UserID   int
	Email    string
	UserType string

	UserProfileID int
	MemberTierID  int
	FirstName     string
	LastName      string
	Address       string
	CreatedAt     time.Time
	Status        string
}
