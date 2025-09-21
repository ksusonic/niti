package models

type DJ struct {
	ID         int      `json:"id"`
	TelegramID *int64   `json:"user_id"`
	StageName  string   `json:"stage_name"`
	AvatarURL  *string  `json:"avatar_url"`
	Socials    []Social `json:"socials"`
}

type Social struct {
	Name string  `json:"name"`
	URL  string  `json:"url"`
	Icon *string `json:"icon"`
}
