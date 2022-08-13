package playlist

import (
	"api/pkg/track"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Playlist struct {
	Name   string        `json:"name"`
	Tracks []track.Track `json:"tracks"`
}

func GetPlaylists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var data map[string]interface{}
	var url string
	var playlists []Playlist
	limit := 0
	offset := 0

	client := &http.Client{}
	accessToken := r.Header.Get("Authorization")
	requestBody, _ := io.ReadAll(r.Body)

	if err := json.Unmarshal(requestBody, &data); err != nil {
		log.Panic(err)
	}

	if data["limit"] != nil{
		limit = int(data["limit"].(float64))
	} 
	if data["offset"] != nil{
		offset = int(data["limit"].(float64))
	}

	log.Println("LIMIT:", limit, "OFFSET:", offset)

	if accessToken == "" {
		http.Error(w, "Missing Parameters", http.StatusBadRequest)
		return
	}

	if limit != 0 || offset != 0 {
		url = fmt.Sprintf("https://api.spotify.com/v1/me/playlists?limit=%d&offset=%d", int(limit), int(offset))
	} else {
		url = "https://api.spotify.com/v1/me/playlists"
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
		name := p.(map[string]interface{})["name"].(string)
		tracksInfo := p.(map[string]interface{})["tracks"].(map[string]interface{})
		tracksHref := tracksInfo["href"].(string)
		
		tracks, err := getPlaylistTracks(accessToken, tracksHref)
		if err != nil {
			log.Panic(err)
			http.Error(w, "can't parse data", http.StatusInternalServerError)
		}
		playlist := Playlist{Name: name, Tracks: tracks}
		playlists = append(playlists, playlist)
		
	}
	result, err := json.Marshal(&playlists)
	if err != nil {
		log.Panic(err)
	}
	log.Println(string(result))
	w.Write(result)
}

func getPlaylistTracks(accessToken string, playlistHref string) ([]track.Track, error) {
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
	tracks, err := track.GetMultipleTracks(trackIds, accessToken)
	if err != nil {
		return nil, err
	}

	return tracks.([]track.Track), nil
}
