package gopushbullet

import (
	"encoding/json"
)

type Subscription struct {
	Active   bool    `json:"active"`
	Channel  Channel `json:"channel"`
	Created  float64 `json:"created"`
	Iden     string  `json:"iden"`
	Modified float64 `json:"modified"`
	Muted    bool    `json:"muted"`
}

type Channel struct {
	Description string `json:"description"`
	Iden        string `json:"iden"`
	ImageURL    string `json:"image_url"`
	Name        string `json:"name"`
	Tag         string `json:"tag"`
}

type SubscriptionResponse struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

type SubscriptionService struct {
	client *Client
}

func NewSubscriptionService(client *Client) *SubscriptionService {
	return &SubscriptionService{client}
}

type SubscriptionListCall struct {
	service *SubscriptionService
}

func (s *SubscriptionService) List() *SubscriptionListCall {
	return &SubscriptionListCall{
		service: s,
	}
}

func (c *SubscriptionListCall) Do() (*[]Subscription, error) {
	data, err := c.service.client.run("GET", "subscriptions", nil)
	if err != nil {
		return nil, err
	}

	var s SubscriptionResponse
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}

	return &s.Subscriptions, nil
}
