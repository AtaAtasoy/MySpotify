package models

type ClientCredentials struct{
	Access_token string
	Token_type string
	Expires_in float64
	Refresh_token string
	User_id string
	Images []interface{}
	Display_name string
}