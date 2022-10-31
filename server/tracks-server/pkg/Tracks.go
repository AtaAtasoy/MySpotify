package tracks

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetTopTracks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	client := &http.Client{}
	requestBody, _ := io.ReadAll(r.Body)
	var data map[string]interface{}
	var tracks []Track
	var url string

	if err := json.Unmarshal(requestBody, &data); err != nil {
		log.Panic(err)
	}
	accessToken := data["access_token"]
	limit := data["limit"]

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
		http.Error(w, res.Status, http.StatusBadRequest)
		return
	}
	defer res.Body.Close()

	responseBody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(responseBody, &data)
	if err != nil {
		log.Panic(err)
		http.Error(w, "can't parse data", http.StatusInternalServerError)
	}

	for _, track := range data["items"].([]interface{}) {
		parsedTrack, err := ParseTrackData(track.(map[string]interface{}), accessToken.(string))
		if err != nil {
			log.Panic(err)
			http.Error(w, "can't parse data", http.StatusInternalServerError)
		}
		tracks = append(tracks, parsedTrack.(Track))
	}
	tracks, err = getAudioAnalysis(tracks, accessToken.(string))
	if err != nil{
		log.Panic(err)
		http.Error(w, "can't get audio analysis", http.StatusInternalServerError)
	}

	result, err := json.Marshal(&tracks)
	if err != nil {
		log.Panic(err)
	}
	w.Write(result)
}

//TODO: Fetch track info concurrently for each list of track ids.
func GetMultipleTracks(trackIds [][]string, accessToken string) (interface{}, error) {
	var data map[string]interface{}
	var parsedTracks []Track
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
		log.Println("Request URL: ", url)

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
			if track != nil {
				parsedTrack, err := ParseTrackData(track.(map[string]interface{}), accessToken)
				if err != nil {
					log.Panic(err)
					return nil, err
				}
				parsedTracks = append(parsedTracks, parsedTrack.(Track))
			}
		}
		parsedTracks, err = getAudioAnalysis(parsedTracks, accessToken)
		if err != nil{
			log.Panic(err)
		}
	}
	return parsedTracks, nil
}

func getAudioAnalysis(parsedTracks []Track, accessToken string) ([]Track, error) {
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
				parsedTracks[i].SetAudioFeatures(feature.(map[string]interface{}))
			}
		}
	}
	return parsedTracks, nil
}