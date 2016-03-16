package gopushbullet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

type clientInfo struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json: "client_secret"`
}

func GetFullClient(s string) *Client {
	client := getClient(s)
	return &Client{client: client}
}

func getClient(s string) *http.Client {

	file, err := os.Open(s)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	var info clientInfo
	err = json.Unmarshal(fileContents, &info)
	if err != nil {
		panic(err)
	}

	conf := &oauth2.Config{
		ClientID:     info.ClientID,
		ClientSecret: info.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.pushbullet.com/authorize",
			TokenURL: "https://api.pushbullet.com/oauth2/token",
		},
		RedirectURL: "http://localhost:3000",
	}

	// First try to Read Token
	token, err := getTokenFromFile("pushbullet-token.json")
	if err != nil {
		token = getTokenFromWeb(conf)
	}

	return conf.Client(oauth2.NoContext, token)
}

func getTokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}

	err = json.NewDecoder(f).Decode(t)
	defer f.Close()

	return t, err
}

func saveTokenToFile(filename string, token *oauth2.Token) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(token)
	if err != nil {
		panic(err)
	}
}

func getTokenFromWeb(conf *oauth2.Config) *oauth2.Token {
	grantType := oauth2.SetAuthURLParam("response_type", "code")

	url := conf.AuthCodeURL("goingtoignorethiskindoffornow", oauth2.AccessTypeOffline, grantType)
	fmt.Printf("Visit the URL for the auth dialog: %v \n", url)

	var code string
	if _, err := fmt.Scan(&code); err != nil {
		panic(err)
	}

	token, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		panic(err)
	}

	saveTokenToFile("pushbullet-token.json", token)

	return token
}
