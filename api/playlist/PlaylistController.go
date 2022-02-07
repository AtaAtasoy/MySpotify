package playlist

import (
	"api/models"
	"api/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetPlaylists(w http.ResponseWriter, r *http.Request){
	var data map[string]interface{}
	var url string
	var playlists []models.Playlist
	
	util.EnableCors(&w)
	client := &http.Client{}
	requestBody, _ := io.ReadAll(r.Body)

	if err := json.Unmarshal(requestBody, &data); err != nil {
		log.Panic(err)
	}
	accessToken := data["access_token"]
	limit := data["limit"]

	if accessToken == nil {
		http.Error(w, "Missing Parameters", http.StatusBadRequest)
		return
	}

	if limit != nil {
		url = fmt.Sprintf("https://api.spotify.com/v1/me/playlists?limit=%s", limit)
	} else {
		url = "https://api.spotify.com/v1/me/playlists"
	}
	log.Println(url)

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

	for _, p := range data["items"].([]interface{}) {
		owner := p.(map[string]interface{})["owner"].(map[string]interface{})["id"]
		log.Println("Owner", owner)
		name := p.(map[string]interface{})["name"].(string)
		tracksInfo :=  p.(map[string]interface{})["tracks"].(map[string]interface{})
		tracksHref := tracksInfo["href"].(string)

		tracks, err := getPlaylistTracks(accessToken.(string), tracksHref)
		if err != nil {
			log.Panic(err)
			http.Error(w, "can't parse data", http.StatusInternalServerError)
		}
		log.Println(tracks)
		playlist := models.Playlist{Name: name, Tracks: tracks}
		playlists = append(playlists, playlist)
	}
	result, err := json.Marshal(&playlists)
	if err != nil {
		log.Panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}


func getPlaylistTracks(accessToken string, playlistHref string) ([]models.Track, error){
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
		trackData := track.(map[string]interface{})["track"].(map[string]interface{})
		ids = append(ids, trackData["id"].(string))
		if len(ids) == 50{
			trackIds = append(trackIds, ids)
			ids = nil
		}
	}
	trackIds = append(trackIds, ids)
	tracks, err := util.GetMultipleTracks(trackIds, accessToken)
	if err != nil{
		return nil, err
	}
	log.Println(tracks)

	return tracks.([]models.Track), nil
}