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

func ParseTrackData(track map[string]interface{}, accessToken string) (interface{}, error) {
	var parsedTrack models.Track
	var artistIds [][]string
	var tempArtistIds []string

	id := track["id"].(string)
	name := track["name"].(string)
	popularity := track["popularity"].(float64)
	for _, artist := range track["artists"].([]interface{}) {
		artistId := artist.(map[string]interface{})["id"].(string)
		tempArtistIds = append(tempArtistIds, artistId)
		if len(tempArtistIds) == 50 {
			artistIds = append(artistIds, tempArtistIds)
			tempArtistIds = nil
		}
	}
	artistIds = append(artistIds, tempArtistIds)
	parsedArtists, err := GetArtistData(accessToken, artistIds)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	parsedTrack = models.Track{Id: id, Name: name, Popularity: popularity, Artists: parsedArtists.([]models.Artist)}
	return parsedTrack, nil
}

func GetMultipleTracks(trackIds [][]string, accessToken string) (interface{}, error) {
	var data map[string]interface{}
	var parsedTracks []models.Track
	client := &http.Client{}
	for _, idList := range trackIds {
		url := "https://api.spotify.com/v1/tracks?ids="
		log.Println("TRACK IDS:", idList, "Length:", len(idList))
		for index, trackId := range idList {
			if index == len(idList) - 1 {
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

	}
	return parsedTracks, nil
}
