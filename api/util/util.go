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

func GetArtistData(accessToken string, artistIds [][]string) (interface{}, error) {
	var data map[string]interface{}
	var parsedArtists []models.Artist
	client := &http.Client{}
	for _, idList := range artistIds {
		url := "https://api.spotify.com/v1/artists?ids="
		for index, artistId := range idList {
			if index == len(idList)-1 {
				url = url + artistId
			} else {
				url = url + artistId + ","
			}
		}
		log.Println("Request: URL", url)
		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Panic(err)
		}
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		request.Header.Add("Content-Type", "application/json")

		res, err := client.Do(request)
		if err != nil {
			log.Panic(err)
		}
		log.Println(res.Status)
		if res.Status != "200 OK" {
			return nil, errors.New(res.Status)
		}
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

		for _, artist := range data["artists"].([]interface{}) {
			parsedArtist := ParseArtistData(artist.(map[string]interface{}))
			if err != nil {
				log.Fatalln(err)
				return nil, err
			}
			parsedArtists = append(parsedArtists, parsedArtist)
		}
	}
	return parsedArtists, nil
}

func ParseArtistData(artist map[string]interface{}) models.Artist {
	return models.Artist{Id: artist["id"].(string), Name: artist["name"].(string), Popularity: artist["popularity"].(float64)}
}

/**
Acousticness float32
	Danceability float32
	Duration_ms float32
	Energy float32
	Instrumentalness float32
	Liveness float32
	Loudness float32
	Mode float32
	Speechiness float32
	Tempo float32
	Valence float32
*/
func ParseTrackData(track map[string]interface{}, accessToken string) (interface{}, error) {
	id := track["id"].(string)
	name := track["name"].(string)
	popularity := track["popularity"].(float64)

	return models.Track{Id: id, Name: name, Popularity: popularity}, nil
}

func GetMultipleTracks(trackIds [][]string, accessToken string) (interface{}, error) {
	var data map[string]interface{}
	var parsedTracks []models.Track
	client := &http.Client{}
	for _, idList := range trackIds {
		if len(idList) == 0{
			continue
		}
		url := "https://api.spotify.com/v1/tracks?ids="
		for index, trackId := range idList {
			if index == len(idList)-1 {
				url = url + trackId
			} else {
				url = url + trackId + ","
			}
		}
		log.Println("Request: URL", url)

		request, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Panic(err)
		}
		request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		request.Header.Add("Content-Type", "application/json")

		res, err := client.Do(request)
		if err != nil {
			log.Panic(err)
		}
		log.Println(res.Status)
		if res.Status != "200 OK" {
			return nil, errors.New(res.Status)
		}
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
		for _, track := range data["tracks"].([]interface{}) {
			parsedTrack, err := ParseTrackData(track.(map[string]interface{}), accessToken)
			if err != nil {
				log.Panic(err)
				return nil, err
			}
			parsedTracks = append(parsedTracks, parsedTrack.(models.Track))
		}
		parsedTracks, err = GetAudioAnalysis(parsedTracks, accessToken)
		if err != nil{
			log.Panic(err)
		}
	}
	return parsedTracks, nil
}

func GetAudioAnalysis(parsedTracks []models.Track, accessToken string) ([]models.Track, error) {
	var data map[string]interface{}
	client := &http.Client{}
	url := "https://api.spotify.com/v1/audio-features?ids="
	for index, track := range parsedTracks {
		if index == len(parsedTracks)-1 {
			url = url + track.Id
		} else {
			url = url + track.Id + ","
		}
	}
	log.Println("Request: URL", url)

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {
		log.Panic(err)
	}
	log.Println(res.Status)
	if res.Status != "200 OK" {
		return nil, errors.New(res.Status)
	}
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
	for _, feature := range data["audio_features"].([]interface{}) {
		for i := range parsedTracks {
			if parsedTracks[i].Id == feature.(map[string]interface{})["id"].(string) {
				parsedTracks[i].Acousticness = feature.(map[string]interface{})["acousticness"].(float64)
				parsedTracks[i].Danceability = feature.(map[string]interface{})["danceability"].(float64)
				parsedTracks[i].Duration_ms = feature.(map[string]interface{})["duration_ms"].(float64)
				parsedTracks[i].Energy = feature.(map[string]interface{})["energy"].(float64)
				parsedTracks[i].Instrumentalness = feature.(map[string]interface{})["instrumentalness"].(float64)
				parsedTracks[i].Liveness = feature.(map[string]interface{})["liveness"].(float64)
				parsedTracks[i].Loudness = feature.(map[string]interface{})["loudness"].(float64)
				parsedTracks[i].Mode = feature.(map[string]interface{})["mode"].(float64)
				parsedTracks[i].Speechiness = feature.(map[string]interface{})["speechiness"].(float64)
				parsedTracks[i].Tempo = feature.(map[string]interface{})["tempo"].(float64)
				parsedTracks[i].Valence = feature.(map[string]interface{})["valence"].(float64)
			}
		}
	}
	return parsedTracks, nil
}
