package authentication

import (
	// "api/models"
	"api/util"
	"encoding/base64"
	"fmt"
	"io/ioutil"
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
func GetAccessToken(w http.ResponseWriter, req *http.Request) {
	util.EnableCors(&w)
	code := req.URL.Query().Get("code")
	state := req.URL.Query().Get("state")

	if len(code) < 1 {
		log.Fatalln("Code is missing")
	} else if len(state) < 16 {
		log.Fatalln("state_mismatch")
	} else {
		// Create the post request to fetch the access token
		client := &http.Client{}
		data := url.Values{}
		data.Set("grant_type", "authorization_code")
		data.Set("redirect_uri", os.Getenv("REDIRECT_URI"))
		data.Set("code", code)

		url := "https://accounts.spotify.com/api/token"

		r, err := http.NewRequest("POST", url, strings.NewReader(data.Encode()))
		if err != nil {
			log.Fatalln(err)
		}

		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		r.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.RawStdEncoding.EncodeToString([]byte(os.Getenv("CLIENT_ID")+":"+os.Getenv("CLIENT_SECRET")))))

		res, err := client.Do(r)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(res.Status)
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		w.Write(body)		
	}
}