package gopushbullet

import (
	"encoding/json"
)

// Example Usage: gear, err := strava.NewGearService(client).Get(gearId).Do()

type User struct {
	Created         float64 `json:"created"`
	Email           string  `json:"email"`
	EmailNormalized string  `json:"email_normalized"`
	Iden            string  `json:"iden"`
	ImageURL        string  `json:"image_url"`
	MaxUploadSize   int     `json:"max_upload_size"`
	Modified        float64 `json:"modified"`
	Name            string  `json:"name"`
}

type UserService struct {
	client *Client
}

type UserGetCall struct {
	service *UserService
	id      string
}

func NewUserService(client *Client) *UserService {
	return &UserService{client}
}

func (s *UserService) Get() *UserGetCall {
	return &UserGetCall{
		service: s,
	}
}

func (c *UserGetCall) Do() (*User, error) {
	data, err := c.service.client.run("GET", "users/me", nil)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
