package util

import (
	"api/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
)

func GenerateRandomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)

	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func GetArtistData(bearerToken string, id string) (interface{}, error){
	var data map[string]interface{}
	client := &http.Client{}
	url := "https://api.spotify.com/v1/artists/" + id
	request, err := http.NewRequest("GET", url, nil)
	if err != nil{
		log.Fatalln(err)
	}
	request.Header.Add("Authorization", "Bearer " + bearerToken)
	request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {		
		log.Fatalln(err)
		return nil, err
	}
	log.Println("Got artist data" + res.Status)
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	
	return models.Artist{Id: id, Name: data["name"].(string), Popularity: data["popularity"].(float64)}, nil
}

func ParseTrackData(track map[string]interface{}, accessToken string) (interface{}, error){
	var artists []models.Artist
	var parsedTrack models.Track
	
	id := track["id"].(string)
	name := track["name"].(string)
	popularity := track["popularity"].(float64)
	for _, artist := range track["artists"].([]interface{}) {
		id := artist.(map[string]interface{})["id"].(string)
		a, err := GetArtistData(accessToken, id)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		artists = append(artists, a.(models.Artist))
	}
	parsedTrack =  models.Track{Id: id, Name: name, Popularity: popularity, Artists: artists}
	
	return parsedTrack, nil
}

func GetMultipleTracks(trackIds [][]string, accessToken string) (interface{}, error){
	var data map[string]interface{}
	var parsedTracks []models.Track
	client := &http.Client{}
	url := "https://api.spotify.com/v1/tracks?ids="
	for _, idList := range(trackIds){
		for index, trackId := range(idList){
			if index == 49{
				url = url + trackId
			} else{
				url = url + trackId + ","
			}
		}

		fmt.Println(url)
		fmt.Println("Here")
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
			return nil, errors.New(res.Status)
		}
		defer res.Body.Close()
	
		responseBody, err := io.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		err = json.Unmarshal(responseBody, &data)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		for _, track := range(data["tracks"].([]interface{})){
			parsedTrack, err := ParseTrackData(track.(map[string]interface{}), accessToken)
			if err != nil {
				log.Fatalln(err)
				return nil, err
			}
			parsedTracks = append(parsedTracks, parsedTrack.(models.Track))
		}
	}
	return parsedTracks, nil
}
	// parsedTrack, err := util.ParseTrackData(trackData, accessToken)
	// log.Println(parsedTrack)
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return nil, err
	// }
	// tracks = append(tracks, parsedTrack.(models.Track))
	
