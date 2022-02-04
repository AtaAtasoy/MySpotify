package main

import (
	"api/artist"
	"api/authentication"
	"api/playlist"
	"api/track"
	"api/user"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	http.HandleFunc("/", authentication.GetUserAuthorization)
	http.HandleFunc("/callback", authentication.GetAccessToken)
	http.HandleFunc("/me", user.GetUserProfile)
	http.HandleFunc("/me/favorite/tracks", track.GetTopTracks)
	http.HandleFunc("/me/favorite/artists", artist.GetTopArtists)
	http.HandleFunc("/me/playlists", playlist.GetPlaylists)
	http.ListenAndServe(":8080", nil)
}