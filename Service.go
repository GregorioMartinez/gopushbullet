package gopushbullet

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const basePath = "https://api.pushbullet.com/"
const version = "v2/"

type Client struct {
	client *http.Client
}

func (client *Client) run(method, path string, params map[string]interface{}) ([]byte, error) {
	var err error

	values := make(url.Values)
	for k, v := range params {
		values.Set(k, fmt.Sprintf("%v", v))
	}

	req, err := http.NewRequest(method, basePath+version+path+"?"+values.Encode(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, err
}
