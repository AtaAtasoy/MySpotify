package main

import (
	"api/pkg/artist"
	"api/pkg/playlist"
	"api/pkg/track"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	godotenv.Load(".env")

	router := mux.NewRouter()
	
	router.HandleFunc("/tracks", track.GetTopTracks)
	router.HandleFunc("/artists", artist.GetTopArtists)
	router.HandleFunc("/playlists", playlist.GetPlaylists)

	//TODO:Setup CORS access origin
	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}