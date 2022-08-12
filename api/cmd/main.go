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

	router.HandleFunc("/me/favorite/tracks", track.GetTopTracks)
	router.HandleFunc("/me/favorite/artists", artist.GetTopArtists)
	router.HandleFunc("/me/playlists", playlist.GetPlaylists)

	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":8080", handler))
}