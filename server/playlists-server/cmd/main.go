package main

import (
	"fmt"
	"log"
	"net/http"
	"playlists-server/playlists"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to playlists-server")
}

func main() {
	godotenv.Load(".env")

	router := mux.NewRouter()

	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/playlists", playlists.GetPlaylists)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*.ngrok.io"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT"},
		AllowedHeaders: []string{"Accept", "Accept-Language", "Authorization", "Content-Type", "Origin", "Referer", "Accept", "User-Agent", "Username",},
	})

	//TODO:Setup CORS access origin
	handler := c.Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
