package models

type NotificationRequest struct {
	EventID          int     `json:"event_id"`
	Message          string  `json:"message"`
	NotificationType string  `json:"notification_type"`
	TargetUsers      []int64 `json:"target_users"`
}

type NotificationResponse struct {
	Success   bool `json:"success"`
	SentCount int  `json:"sent_count"`
}
