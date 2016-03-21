package gopushbullet

import (
	"golang.org/x/oauth2"
)

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://www.pushbullet.com/authorize",
	TokenURL: "https://api.pushbullet.com/oauth2/token",
}
