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
func GetUserProfile(w http.ResponseWriter, req *http.Request){
	util.EnableCors(&w)

	client := &http.Client{}
	requestBody, _ := io.ReadAll(req.Body)
	var data map[string]string
	if err := json.Unmarshal(requestBody, &data); err != nil{
		log.Fatalln(err)
	}
	access_token := data["access_token"]
	url := "https://api.spotify.com/v1/me"

	r, err := http.NewRequest("GET", url, nil)
	if err != nil{
		log.Fatalln(err)
	}
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", access_token))
	r.Header.Add("Content-Type", "application/json")

	res, err := client.Do(r, )
	if err != nil{
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

func GetUserTopItems(w http.ResponseWriter, req *http.Request){
	util.EnableCors(&w)
	client := &http.Client{}
	request_body, _ := io.ReadAll(req.Body)
	var data map[string]string

	if err := json.Unmarshal(request_body, &data); err != nil{
		log.Fatalln(err)
	}
	access_token := data["access_token"]
	item_type := data["item_type"]

	url := fmt.Sprintf("https://api.spotify.com/v1/me/top/%s", item_type)

	r, err := http.NewRequest("GET", url, nil)
	if err != nil{
		log.Fatalln(err)
	}
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", access_token))
	r.Header.Add("Content-Type", "application/json")

	res, err := client.Do(r, )
	if err != nil{
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