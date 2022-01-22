package main

import(
	"api/authentication"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	http.HandleFunc("/", authentication.GetUserAuthorization)
	http.HandleFunc("/callback", authentication.GetAccessToken)
	http.ListenAndServe(":8080", nil)
}


