package entity

import "time"

type UserProfile struct {
	UserProfileID int
	UserId        int
	MemberTierId  int
	FirstName     string
	LastName      string
	Address       string
	CreatedAt     time.Time
	Status        string
}
