package models

type AccessTokenResponse struct{
	Access_token string
	Token_type string
	Scope string
	Expires_in int
	Refresh_token string
}