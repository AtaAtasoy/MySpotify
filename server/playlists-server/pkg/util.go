package playlists
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	t "tracks-server/tracks"
	"github.com/shopspring/decimal"
)

type Playlist struct {
	Id string `json:"id"`
	Name   string        `json:"name"`
	Tracks []t.Track `json:"tracks"`
	Images []interface{} `json:"images"`
	Attributes map[string]float64 `json:"attributes"`
}

func FetchPlaylists(accessToken string, username string, url string) ([]Playlist, error) {
	var data map[string]interface{}
	var playlists []Playlist
	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(res.Status)
	if res.Status != "200 OK" {
		return nil, err
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

	for _, p := range data["items"].([]interface{}) {
		ownerData :=  p.(map[string]interface{})["owner"].((map[string]interface{}))
		if isOwnedByLoggedInUser(ownerData, username) {
			id :=  p.(map[string]interface{})["id"].(string)
			name := p.(map[string]interface{})["name"].(string)
			images := p.(map[string]interface{})["images"].([]interface{})
			tracksInfo := p.(map[string]interface{})["tracks"].(map[string]interface{})
			tracksHref := tracksInfo["href"].(string)

			tracks, err := getPlaylistTracks(accessToken, tracksHref)
			if err != nil {
				log.Panic(err)
				return nil, err
			}
			playlist := Playlist{Id: id, Name: name, Tracks: tracks, Images: images, Attributes: calculatePlaylistAttributes(tracks)}
			playlists = append(playlists, playlist)
		}
	}
	return playlists, nil
}

func getPlaylistTracks(accessToken string, playlistHref string) ([]t.Track, error) {
	var data map[string]interface{}
	var trackIds [][]string
	var ids []string

	client := &http.Client{}
	request, err := http.NewRequest("GET", playlistHref, nil)
	if err != nil {
		log.Panic(err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {
		log.Panic(err)
		return nil, err
	}
	log.Println(res.Status)
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

	for _, track := range data["items"].([]interface{}) {
		if track != nil {
			trackData := track.(map[string]interface{})["track"]
			if trackData != nil{
				if trackData.(map[string]interface{})["id"] != nil{
					ids = append(ids, trackData.(map[string]interface{})["id"].(string))
					if len(ids) == 50 {
						trackIds = append(trackIds, ids)
						ids = nil
					}
				}
			}
		}
	}
	trackIds = append(trackIds, ids)
	tracks, err := t.GetMultipleTracks(trackIds, accessToken)
	if err != nil {
		return nil, err
	}

	return tracks.([]t.Track), nil
}

func calculatePlaylistAttributes(tracks []t.Track) map[string]float64{
	length := len(tracks)
	popularity := 0.0
	acousticness := 0.0
	danceability := 0.0
	duration := 0.0
	energy := 0.0
	instrumentalness := 0.0
	liveness := 0.0
	speechiness := 0.0
	tempo := 0.0
	valence := 0.0

	for _, track := range(tracks){
		popularity += track.Popularity
		acousticness += track.Acousticness
		danceability += track.Danceability
		duration += track.Duration_ms
		energy += track.Energy
		instrumentalness += track.Instrumentalness
		liveness += track.Liveness
		speechiness += track.Speechiness
		tempo += track.Tempo
		valence += track.Valence
	}
	attributes := make(map[string]float64)

	attributes["popularity"], _ = decimal.NewFromFloatWithExponent( (popularity / float64(length) ), -2).Float64()
	attributes["acousticness"], _ = decimal.NewFromFloatWithExponent( (acousticness / float64(length) ) * 100, -2).Float64()
	attributes["danceability"], _ =  decimal.NewFromFloatWithExponent( (danceability / float64(length) ) * 100, -2).Float64()
	attributes["energy"], _ =  decimal.NewFromFloatWithExponent( (energy / float64(length) ) * 100, -2).Float64()
	attributes["instrumentalness"], _ =  decimal.NewFromFloatWithExponent( (instrumentalness / float64(length) ) * 100, -2).Float64()
	attributes["speechiness"], _ =  decimal.NewFromFloatWithExponent( (speechiness / float64(length) ) * 100, -2).Float64()
	attributes["valence"], _ =  decimal.NewFromFloatWithExponent( (valence / float64(length) ) * 100, -2).Float64()
	attributes["duration"], _ = decimal.NewFromFloatWithExponent( (duration / float64(length) ) / 1000, -2).Float64()
	attributes["tempo"], _ = decimal.NewFromFloatWithExponent( (tempo / float64(length) ), -2).Float64()

	return attributes
}

func isOwnedByLoggedInUser(ownerData map[string]interface{}, username string) bool{
	return ownerData["id"].(string) == username
}