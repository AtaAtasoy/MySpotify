package user

import (
	"api/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Receives Spotify access token from the request header
// Fetches the information about the user
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)

	client := &http.Client{}
	requestBody, _ := io.ReadAll(r.Body)
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