package gopushbullet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const basePath = "https://api.pushbullet.com/"
const version = "v2/"

type Client struct {
	Client *http.Client
}

func (client *Client) run(method, path string, params map[string]interface{}) ([]byte, error) {
	var err error

	values := make(url.Values)
	for k, v := range params {
		values.Set(k, fmt.Sprintf("%v", v))
	}

	var req *http.Request
	if method == "POST" {
		j, err := json.Marshal(params)
		if err != nil {
			panic(err)
		}

		r := bytes.NewBuffer(j)
		req, err = http.NewRequest("POST", basePath+version+path, r)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")

	} else {
		req, err = http.NewRequest(method, basePath+version+path+"?"+values.Encode(), nil)
		if err != nil {
			return nil, err
		}
	}

	resp, err := client.Client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	/*
		200 OK - Everything worked as expected.
		400 Bad Request - Usually this results from missing a required parameter.
		401 Unauthorized - No valid access token provided.
		403 Forbidden - The access token is not valid for that request.
		404 Not Found - The requested item doesn't exist.
		429 Too Many Requests - You have been ratelimited for making too many requests to the server.
		5XX Server Error
	*/

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var err ErrorResp
		json.Unmarshal(body, &err)
		return nil, err.Error
	}

	return body, err
}
