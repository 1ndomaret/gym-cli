package entity

import "time"

type Event struct {
	EventID     int
	EventName   string
	MinTierRank int
	Schedule    time.Time
}
