package user

import (
	"api/models"
	"api/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Receives Spotify access token from the request header
// Fetches the information about the user
func GetUserProfile(w http.ResponseWriter, req *http.Request) {
	util.EnableCors(&w)

	client := &http.Client{}
	requestBody, _ := io.ReadAll(req.Body)
	var data map[string]string
	if err := json.Unmarshal(requestBody, &data); err != nil {
		log.Fatalln(err)
	}
	accessToken := data["access_token"]
	url := "https://api.spotify.com/v1/me"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Status)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(body)
}

func GetUserTopItems(w http.ResponseWriter, req *http.Request) {
	util.EnableCors(&w)
	client := &http.Client{}
	requestBody, _ := io.ReadAll(req.Body)
	var data map[string]interface{}
	var tracks []models.Track
	var artists []models.Artist

	if err := json.Unmarshal(requestBody, &data); err != nil {
		log.Fatalln(err)
	}
	accessToken := data["access_token"]
	itemType := data["item_type"]
	limit := data["limit"]

	url := fmt.Sprintf("https://api.spotify.com/v1/me/top/%s?limit=%s", itemType, limit)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
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
		log.Fatalln(err)
	}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		log.Fatalln(err)
	}

	if itemType == "tracks" {
		for _, track := range data["items"].([]interface{}) {
			id := track.(map[string]interface{})["id"].(string)
			name := track.(map[string]interface{})["name"].(string)
			popularity := track.(map[string]interface{})["popularity"].(float64)
			for _, artist := range track.(map[string]interface{})["artists"].([]interface{}) {
				id := artist.(map[string]interface{})["id"].(string)
				name := artist.(map[string]interface{})["name"].(string)
				uri := artist.(map[string]interface{})["uri"].(string)

				artists = append(artists, models.Artist{Id: id, Name: name, Uri: uri})
			}
			tracks = append(tracks, models.Track{Id: id, Name: name, Popularity: popularity, Artists: artists})
		}
		result, err := json.Marshal(&tracks)
		if err != nil {
			log.Fatalln(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
		return
	} else if itemType == "artists" {
		for _, artist := range data["items"].([]interface{}) {
			id := artist.(map[string]interface{})["id"].(string)
			name := artist.(map[string]interface{})["name"].(string)
			uri := artist.(map[string]interface{})["uri"].(string)

			artists = append(artists, models.Artist{Id: id, Name: name, Uri: uri})
		}
		result, err := json.Marshal(&artists)
		if err != nil {
			log.Fatalln(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Item Type!"))
	}
}
