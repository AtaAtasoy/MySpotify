package track

import (
	"api/models"
	"api/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetTopTracks(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	client := &http.Client{}
	requestBody, _ := io.ReadAll(r.Body)
	var data map[string]interface{}
	var tracks []models.Track
	var artists []models.Artist
	var url string

	if err := json.Unmarshal(requestBody, &data); err != nil {
		log.Fatalln(err)
	}
	accessToken := data["access_token"]
	limit := data["limit"]
	log.Println(limit)

	if accessToken == nil {
		http.Error(w, "Missing Parameters", http.StatusBadRequest)
		return
	}

	if limit != nil {
		url = fmt.Sprintf("https://api.spotify.com/v1/me/top/tracks?limit=%s", limit)
	} else {
		url = "https://api.spotify.com/v1/me/top/tracks"
	}
	log.Println(url)

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
		http.Error(w, "can't parse data", http.StatusInternalServerError)
	}

	for _, track := range data["items"].([]interface{}) {
		id := track.(map[string]interface{})["id"].(string)
		name := track.(map[string]interface{})["name"].(string)
		popularity := track.(map[string]interface{})["popularity"].(float64)
		for _, artist := range track.(map[string]interface{})["artists"].([]interface{}) {
			id := artist.(map[string]interface{})["id"].(string)
			a, err := util.GetArtistData(accessToken.(string), id)
			if err != nil {
				log.Fatalln(err)
				http.Error(w, "can't get artist info", http.StatusInternalServerError)
			}
			artists = append(artists, a.(models.Artist))
		}
		tracks = append(tracks, models.Track{Id: id, Name: name, Popularity: popularity, Artists: artists})
		artists = nil
	}

	result, err := json.Marshal(&tracks)
	if err != nil {
		log.Fatalln(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}