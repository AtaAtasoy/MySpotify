package main

import (
	"api/pkg/artist"
	"api/pkg/authentication"
	"api/pkg/playlist"
	"api/pkg/track"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	router := mux.NewRouter()
	router.Use(accessControlMiddleware)

	router.HandleFunc("/", authentication.GetUserAuthorization)
	router.HandleFunc("/callback", authentication.GetAccessToken)
	router.HandleFunc("/me/favorite/tracks", track.GetTopTracks).Methods("GET")
	router.HandleFunc("/me/favorite/artists", artist.GetTopArtists).Methods("GET")
	router.HandleFunc("/me/playlists", playlist.GetPlaylists).Methods("GET")
	http.ListenAndServe(":8080", handlers.CORS()(router))
}

func accessControlMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
