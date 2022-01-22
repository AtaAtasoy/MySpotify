package main

import(
	"api/authentication"
	"api/user"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	http.HandleFunc("/", authentication.GetUserAuthorization)
	http.HandleFunc("/callback", authentication.GetAccessToken)
	http.HandleFunc("/me", user.GetUserProfile)
	http.HandleFunc("/me/favorite", user.GetUserTopItems)
	http.ListenAndServe(":8080", nil)
}