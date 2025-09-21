package models

import "time"

type Subscription struct {
	UserID    int64
	EventID   int
	CreatedAt time.Time
}
