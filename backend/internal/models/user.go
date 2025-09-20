package models

type User struct {
	TelegramID int64
	Username   string
	FirstName  string
	LastName   *string
	AvatarURL  *string
	IsDJ       bool
}
