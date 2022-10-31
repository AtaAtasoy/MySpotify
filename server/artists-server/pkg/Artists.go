package artists

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Artist struct {
	Id         string        `json:"id"`
	Name       string        `json:"name"`
	Popularity float64       `json:"popularity"`
	Images     []interface{} `json:"images"`
	Genres	[]string `json:"genres"`
}

func GetTopArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	client := &http.Client{}
	var data map[string]interface{}
	var artists []Artist
	var url string

	query := r.URL.Query()
	limit := query["limit"]

	accessToken := r.Header.Get("Authorization")
	log.Print("Received token:", accessToken)
	log.Print("Received limit:", limit)

	if limit != nil {
		url = fmt.Sprintf("https://api.spotify.com/v1/me/top/artists?limit=%s", limit[0])
	} else {
		url = "https://api.spotify.com/v1/me/top/artists"
	}

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

	for _, artist := range data["items"].([]interface{}) {
		parsedArtist := ParseArtistData(artist.(map[string]interface{}))
		artists = append(artists, parsedArtist)
	}

	result, err := json.Marshal(&artists)
	if err != nil {
		log.Fatalln(err)
	}
	w.Write(result)
}

func GetArtistData(accessToken string, artistIds [][]string) (interface{}, error) {
	var data map[string]interface{}
	var parsedArtists []Artist
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

func ParseArtistData(artist map[string]interface{}) Artist {
	return Artist{Id: artist["id"].(string), Name: artist["name"].(string), Popularity: artist["popularity"].(float64), Images: artist["images"].([]interface{}), Genres: artist["genres"].([]string)}
}
