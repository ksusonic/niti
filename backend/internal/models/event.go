package models

import "time"

type Event struct {
	ID                int
	Title             string
	Description       string
	Location          string
	VideoURL          string
	StartsAt          time.Time
	ParticipantsCount int
	IsSubscribed      bool
	DJs               []DJ
}
