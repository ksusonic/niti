package models

import "time"

type Event struct {
	ID          int
	Title       string
	Description string
	Location    *string
	CoverURL    *string
	VideoURL    *string
	StartsAt    *time.Time
	CreatedBy   *int64
}

type EventEnriched struct {
	Event

	ParticipantsCount *int
	IsSubscribed      bool
	DJs               []DJ
}
