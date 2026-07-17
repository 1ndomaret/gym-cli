package entity

import "time"

type UserDetail struct {
	UserId        int
	UserProfileId int
	UserType      string
	Email         string
	FirstName     string
	LastName      string
	TierName      string
	Address       string
	CreatedAt     time.Time
	Status        bool
}
