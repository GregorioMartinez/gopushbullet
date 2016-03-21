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
	Active          bool    `json:"active"`
	Created         float64 `json:"created"`
	Description     string  `json:"description"`
	Iden            string  `json:"iden"`
	ImageURL        string  `json:"image_url"`
	Modified        float64 `json:"modified"`
	Name            string  `json:"name"`
	SubscriberCount float64 `json:"subscriber_count"`
	Tag             string  `json:"tag"`
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

type SubscriptionCreateCall struct {
	service *SubscriptionService
	args    map[string]interface{}
}

func (s *SubscriptionService) Create() *SubscriptionCreateCall {
	return &SubscriptionCreateCall{
		service: s,
		args:    make(map[string]interface{}),
	}
}

func (c *SubscriptionCreateCall) ChannelTag(tag string) *SubscriptionCreateCall {
	c.args["channel_tag"] = tag
	return c
}

func (c *SubscriptionCreateCall) Do() (*Subscription, error) {
	data, err := c.service.client.run("POST", "subscriptions", c.args)
	if err != nil {
		return nil, err
	}

	var s Subscription
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

type SubscriptionUpdateCall struct {
	service *SubscriptionService
	iden    string
	args    map[string]interface{}
}

func (s *SubscriptionService) Update(iden string) *SubscriptionUpdateCall {
	return &SubscriptionUpdateCall{
		service: s,
		iden:    iden,
		args:    make(map[string]interface{}),
	}
}

func (c *SubscriptionUpdateCall) Muted(muted bool) *SubscriptionUpdateCall {
	c.args["muted"] = muted
	return c
}

func (c *SubscriptionUpdateCall) Do() (*Subscription, error) {
	data, err := c.service.client.run("POST", "subscriptions/"+c.iden, c.args)
	if err != nil {
		return nil, err
	}

	var s Subscription
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

type SubscriptionDeleteCall struct {
	service *SubscriptionService
	iden    string
}

func (s *SubscriptionService) Delete(iden string) *SubscriptionDeleteCall {
	return &SubscriptionDeleteCall{
		service: s,
		iden:    iden,
	}
}

func (c *SubscriptionDeleteCall) Do() error {
	_, err := c.service.client.run("DELETE", "subscriptions/"+c.iden, nil)
	if err != nil {
		return err
	}
	return nil
}

type SubscriptionChannelInfoCall struct {
	service *SubscriptionService
	args    map[string]interface{}
}

func (s *SubscriptionService) ChannelInfo() *SubscriptionChannelInfoCall {
	return &SubscriptionChannelInfoCall{
		service: s,
		args:    make(map[string]interface{}),
	}
}

func (c *SubscriptionChannelInfoCall) Tag(tag string) *SubscriptionChannelInfoCall {
	c.args["tag"] = tag
	return c
}

func (c *SubscriptionChannelInfoCall) NoRecentPushes(noRecentPushes bool) *SubscriptionChannelInfoCall {
	c.args["no_recent_pushes"] = noRecentPushes
	return c
}

func (c *SubscriptionChannelInfoCall) Do() (*Channel, error) {
	data, err := c.service.client.run("GET", "channel-info", c.args)
	if err != nil {
		return nil, err
	}

	var i Channel
	err = json.Unmarshal(data, &i)
	if err != nil {
		return nil, err
	}

	return &i, nil
}
