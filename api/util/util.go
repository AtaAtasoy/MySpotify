package util

import (
	"api/models"
	"encoding/json"
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