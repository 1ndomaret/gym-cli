package entity

import "time"

type UserDetails struct {
	UserId   int
	Email    string
	UserType string

	UserProfileId int
	MemberTierId  int
	FirstName     string
	LastName      string
	Address       string
	CreatedAt     time.Time
	Status        string
}
