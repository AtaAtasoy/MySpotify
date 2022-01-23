package artist
import (
	"api/models"
	"api/util"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetTopArtists(w http.ResponseWriter, r *http.Request) {
	util.EnableCors(&w)
	client := &http.Client{}
	requestBody, _ := io.ReadAll(r.Body)
	var data map[string]interface{}
	var artists []models.Artist
	var url string

	if err := json.Unmarshal(requestBody, &data); err != nil {
		log.Fatalln(err)
	}
	accessToken := data["access_token"]
	limit := data["limit"]

	if accessToken == nil {
		http.Error(w, "Missing Parameters", http.StatusBadRequest)
		return
	}

	if limit != nil {
		url = fmt.Sprintf("https://api.spotify.com/v1/me/top/artists?limit=%s",limit)
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
		return
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(responseBody))

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, artist := range data["items"].([]interface{}) { 
		id := artist.(map[string]interface{})["id"].(string)
		name := artist.(map[string]interface{})["name"].(string)
		popularity := artist.(map[string]interface{})["popularity"].(float64)

		artists = append(artists, models.Artist{Id: id, Name: name, Popularity: popularity})
	}

	result, err := json.Marshal(&artists)
	if err != nil {
		log.Fatalln(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}