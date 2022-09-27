package playlists

import (
	t "tracks-server/tracks"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Playlist struct {
	Id string `json:"id"`
	Name   string        `json:"name"`
	Tracks []t.Track `json:"tracks"`
	Images []interface{} `json:"images"`
}

func GetPlaylists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data map[string]interface{}
	var url string
	var playlists []Playlist
	var limit string
	client := &http.Client{}

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

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Status)
	if res.Status != "200 OK" {
		http.Error(w, res.Status, http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		log.Panic(err)
		http.Error(w, "can't parse data", http.StatusInternalServerError)
	}
	w.Header().Add("Offset", fmt.Sprintf("%d", int(data["offset"].(float64))))
	w.Header().Add("Limit", fmt.Sprintf("%d", int(data["limit"].(float64))))
	w.Header().Add("Total", fmt.Sprintf("%d", int(data["total"].(float64))))

	for _, p := range data["items"].([]interface{}) {
		ownerData :=  p.(map[string]interface{})["owner"].((map[string]interface{}))
		if isOwnedByLoggedInUser(ownerData, username) {
			id :=  p.(map[string]interface{})["id"].(string)
			name := p.(map[string]interface{})["name"].(string)
			images := p.(map[string]interface{})["images"].([]interface{})
			tracksInfo := p.(map[string]interface{})["tracks"].(map[string]interface{})
			tracksHref := tracksInfo["href"].(string)

			tracks, err := getPlaylistTracks(accessToken, tracksHref)
			if err != nil {
				log.Panic(err)
				http.Error(w, "can't parse data", http.StatusInternalServerError)
			}
			playlist := Playlist{Id: id, Name: name, Tracks: tracks, Images: images}
			playlists = append(playlists, playlist)
		}

	}
	result, err := json.Marshal(&playlists)
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(result))
	w.Write(result)
}

func getPlaylistTracks(accessToken string, playlistHref string) ([]t.Track, error) {
	var data map[string]interface{}
	var trackIds [][]string
	var ids []string

	client := &http.Client{}
	request, err := http.NewRequest("GET", playlistHref, nil)
	if err != nil {
		log.Panic(err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	log.Println(res.Status)
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	for _, track := range data["items"].([]interface{}) {
		if track != nil {
			trackData := track.(map[string]interface{})["track"]
			if trackData != nil{
				if trackData.(map[string]interface{})["id"] != nil{
					ids = append(ids, trackData.(map[string]interface{})["id"].(string))
					if len(ids) == 50 {
						trackIds = append(trackIds, ids)
						ids = nil
					}
				}
			}
		}
	}
	trackIds = append(trackIds, ids)
	tracks, err := t.GetMultipleTracks(trackIds, accessToken)
	if err != nil {
		return nil, err
	}

	return tracks.([]t.Track), nil
}

func isOwnedByLoggedInUser(ownerData map[string]interface{}, username string) bool{
	return ownerData["id"].(string) == username
}