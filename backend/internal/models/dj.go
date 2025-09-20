package models

type DJ struct {
	StageName string
	AvatarURL string
	Socials   []Social
}

type Social struct {
	Name string
	URL  string
	Icon string
}
