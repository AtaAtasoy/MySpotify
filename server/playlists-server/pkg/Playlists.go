package playlists

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func GetPlaylists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var url string
	var limit string

	accessToken := r.Header.Get("Authorization")
	username := r.Header.Get("Username")
	query := r.URL.Query()

	if query["limit"] != nil{
		limit = query.Get("limit")
	} else {
		limit = "50"
	}

	offset := strings.Join(query["offset"], "")

	log.Print("Received token:", accessToken)
	log.Print("Received limit:", limit)
	log.Print("Received offset:", offset)
	log.Print("Received username:", username)

	if accessToken == "" {
		http.Error(w, "Missing Parameters: Authorization", http.StatusBadRequest)
		return
	}

	if offset != "" {
		url = fmt.Sprintf("https://api.spotify.com/v1/me/playlists?limit=50&offset=%s", offset)
	} else {
		url = "https://api.spotify.com/v1/me/playlists?limit=50"
	}
	log.Println("Request URL:", url)

	playlists, err := FetchPlaylists(accessToken, username, url)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	result, err := json.Marshal(&playlists)
	if err != nil {
		log.Panic(err)
	}
	w.Write(result)
}