package models

type JWTAuth struct {
	AccessToken  string
	RefreshToken string
	JTI          string
}
