package gopushbullet

import (
	"encoding/json"
)

type Push struct {
	Active                  bool    `json:"active"`
	Body                    string  `json:"body"`
	Created                 float64 `json:"created"`
	Direction               string  `json:"direction"`
	Dismissed               bool    `json:"dismissed"`
	Iden                    string  `json:"iden"`
	Modified                float64 `json:"modified"`
	ReceiverEmail           string  `json:"receiver_email"`
	ReceiverEmailNormalized string  `json:"receiver_email_normalized"`
	ReceiverIden            string  `json:"receiver_iden"`
	SenderEmail             string  `json:"sender_email"`
	SenderEmailNormalized   string  `json:"sender_email_normalized"`
	SenderIden              string  `json:"sender_iden"`
	SenderName              string  `json:"sender_name"`
	Title                   string  `json:"title"`
	Type                    string  `json:"type"`
}

type PushResponse struct {
	Pushes []Push `json:"pushes"`
}

type PushService struct {
	client *Client
}

type PushListCall struct {
	service *PushService
}

func NewPushService(client *Client) *PushService {
	return &PushService{client}
}

func (s *PushService) List() *PushListCall {
	return &PushListCall{
		service: s,
	}
}

func (c *PushListCall) Do() (*[]Push, error) {
	data, err := c.service.client.run("GET", "pushes", nil)
	if err != nil {
		return nil, err
	}

	var p PushResponse
	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}

	return &p.Pushes, nil
}
