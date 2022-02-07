package authentication

import (
	"api/models"
	"api/util"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Requesting User Authorization
func GetUserAuthorization(w http.ResponseWriter, req *http.Request) {
	util.EnableCors(&w)
	client_id := os.Getenv("CLIENT_ID")
	redirect_uri := os.Getenv("REDIRECT_URI")
	scope := os.Getenv("SCOPE")
	state := util.GenerateRandomString(16)

	redirectUrl := fmt.Sprintf("https://accounts.spotify.com/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s&state=%s", client_id, scope, redirect_uri, state)
	http.Redirect(w, req, redirectUrl, http.StatusSeeOther)
}

// Requesting Access Token
func GetAccessToken(w http.ResponseWriter, r *http.Request) {
	var clientCredentials models.ClientCredentials
	util.EnableCors(&w)
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if len(code) < 1 {
		log.Panic("Code is missing")
		http.Error(w, "Code is missing", http.StatusBadRequest)
		return
	} else if len(state) < 16 {
		log.Panic("state_mismatch")
		http.Error(w, "Code is missing", http.StatusBadRequest)
		return
	} else {
		// Create the post request to fetch the access token
		client := &http.Client{}
		data := url.Values{}
		data.Set("grant_type", "authorization_code")
		data.Set("redirect_uri", os.Getenv("REDIRECT_URI"))
		data.Set("code", code)

		url := "https://accounts.spotify.com/api/token"

		req, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
		if err != nil {
			log.Fatalln(err)
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.RawStdEncoding.EncodeToString([]byte(os.Getenv("CLIENT_ID")+":"+os.Getenv("CLIENT_SECRET")))))

		res, err := client.Do(req)
		if err != nil {
			log.Panic(err)
			http.Error(w, "can't authenticate with spotify", http.StatusBadRequest)
			return
		}
		log.Println(res.Status)
		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Panic(err)
		}

		err = json.Unmarshal(body, &clientCredentials)
		if err != nil{
			log.Panic(err)
			http.Error(w, "can't parse spotify authentication data", http.StatusInternalServerError)
			return
		}
		completeClientCredentials(&clientCredentials)
		result, err := json.Marshal(clientCredentials)
		if err != nil{
			log.Panic(err)
			http.Error(w, "can't parse spotify authentication data", http.StatusInternalServerError)
			return
		}
		w.Write(result)
	}
}

func completeClientCredentials(credentials *models.ClientCredentials) (models.ClientCredentials, error){
	client := &http.Client{}
	var data map[string]interface{}
	url := "https://api.spotify.com/v1/me"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err)
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", credentials.Access_token))
	request.Header.Add("Content-Type", "application/json")

	res, err := client.Do(request)
	if err != nil {
		log.Panic(err)
		return *credentials, err
	}
	log.Println(res.Status)
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
		return *credentials, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil{
		log.Panic(err)
		return *credentials, err
	}

	credentials.Display_name = data["display_name"].(string)
	credentials.Images = data["images"].([]interface{})
	credentials.User_id = data["id"].(string)
	return *credentials, nil
}