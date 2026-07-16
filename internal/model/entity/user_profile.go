package entity

import "time"

type UserProfile struct {
	UserProfileId int
	UserId        int
	MemberTierId  int
	FirstName     string
	LastName      string
	Address       string
	CreatedAt     time.Time
	Status        bool
}
