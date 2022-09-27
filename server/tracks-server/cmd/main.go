package main

import (
	"tracks-server/tracks"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	godotenv.Load(".env")

	router := mux.NewRouter()

	router.HandleFunc("/tracks", tracks.GetTopTracks)

	//TODO:Setup CORS access origin
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}