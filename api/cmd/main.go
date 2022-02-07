package main

import (
	"api/pkg/artist"
	"api/pkg/authentication"
	"api/pkg/playlist"
	"api/pkg/track"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	http.HandleFunc("/", authentication.GetUserAuthorization)
	http.HandleFunc("/callback", authentication.GetAccessToken)
	http.HandleFunc("/me/favorite/tracks", track.GetTopTracks)
	http.HandleFunc("/me/favorite/artists", artist.GetTopArtists)
	http.HandleFunc("/me/playlists", playlist.GetPlaylists)
	http.ListenAndServe(":8080", nil)
}