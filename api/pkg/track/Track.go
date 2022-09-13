package track

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Track struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Popularity float64 `json:"popularity"`
	Acousticness float64 `json:"acousticness"`
	Danceability float64 `json:"danceability"`
	Duration_ms float64 `json:"duration_ms"`
	Energy float64 `json:"energy"`
	Instrumentalness float64 `json:"instrumentalness"`
	Liveness float64 `json:"liveness"`
	Loudness float64 `json:"loudness"`
	Mode float64 `json:"mode"`
	Speechiness float64 `json:"speechiness"`
	Tempo float64 `json:"tempo"`
	Valence float64 `json:"valence"`
}

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
		parsedTrack, err := parseTrackData(track.(map[string]interface{}), accessToken.(string))
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
		fmt.Println("Multiple Tracks Reponse:", data)
		for _, track := range data["tracks"].([]interface{}) {
			if track != nil {
				parsedTrack, err := parseTrackData(track.(map[string]interface{}), accessToken)
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
				parsedTracks[i].setAudioFeatures(feature.(map[string]interface{}))
			}
		}
	}
	return parsedTracks, nil
}

func parseTrackData(track map[string]interface{}, accessToken string) (interface{}, error) {
	id := track["id"].(string)
	name := track["name"].(string)
	popularity := track["popularity"].(float64)

	return Track{Id: id, Name: name, Popularity: popularity}, nil
}

func (t *Track) setAudioFeatures(features map[string]interface{}) {
	t.Acousticness = features["acousticness"].(float64)
	t.Danceability = features["danceability"].(float64)
	t.Duration_ms = features["duration_ms"].(float64)
	t.Energy = features["energy"].(float64)
	t.Instrumentalness = features["instrumentalness"].(float64)
	t.Liveness = features["liveness"].(float64)
	t.Loudness = features["loudness"].(float64)
	t.Mode = features["mode"].(float64)
	t.Speechiness = features["speechiness"].(float64)
	t.Tempo = features["tempo"].(float64)
	t.Valence = features["valence"].(float64)
}