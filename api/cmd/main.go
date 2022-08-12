package main

import (
	"api/pkg/artist"
	"api/pkg/authentication"
	"api/pkg/playlist"
	"api/pkg/track"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	router := mux.NewRouter()
	
	router.HandleFunc("/", authentication.GetUserAuthorization)
	router.HandleFunc("/callback", authentication.GetAccessToken)
	router.HandleFunc("/me/favorite/tracks", track.GetTopTracks)
	router.HandleFunc("/me/favorite/artists", artist.GetTopArtists)
	router.HandleFunc("/me/playlists", playlist.GetPlaylists)

	http.ListenAndServe(":8080", handlers.CORS()(router))
}