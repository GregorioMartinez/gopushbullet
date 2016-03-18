package gopushbullet

import (
	"encoding/json"
)

type Chat struct {
	Active   bool    `json:"active"`
	Created  float64 `json:"created"`
	Iden     string  `json:"iden"`
	Modified float64 `json:"modified"`
	With     With    `json:"with"`
}

type ChatResponse struct {
	Chats []Chat `json:"chats"`
}

type With struct {
	Email           string `json:"email"`
	EmailNormalized string `json:"email_normalized"`
	Iden            string `json:"iden"`
	ImageURL        string `json:"image_url"`
	Name            string `json:"name"`
	Type            string `json:"type"`
}

type ChatService struct {
	client *Client
}

func NewChatService(client *Client) *ChatService {
	return &ChatService{client}
}

type ChatListCall struct {
	service *ChatService
}

func (s *ChatService) List() *ChatListCall {
	return &ChatListCall{
		service: s,
	}
}

func (c *ChatListCall) Do() (*[]Chat, error) {
	data, err := c.service.client.run("GET", "chats", nil)
	if err != nil {
		return nil, err
	}

	var chats ChatResponse
	err = json.Unmarshal(data, &chats)
	if err != nil {
		return nil, err
	}

	return &chats.Chats, nil
}

type ChatCreateCall struct {
	service *ChatService
	args    map[string]interface{}
}

func (s *ChatService) Create(emailaddr string) *ChatCreateCall {
	call := &ChatCreateCall{
		service: s,
		args:    make(map[string]interface{}),
	}

	call.args["email"] = emailaddr

	return call
}

func (c *ChatCreateCall) Do() (*Chat, error) {
	data, err := c.service.client.run("POST", "chats", c.args)
	if err != nil {
		return nil, err
	}

	var chat Chat
	err = json.Unmarshal(data, &chat)
	if err != nil {
		return nil, err
	}

	return &chat, nil
}

type ChatUpdateCall struct {
	service *ChatService
	iden    string
	args    map[string]interface{}
}

func (s *ChatService) Update(iden string) *ChatUpdateCall {
	return &ChatUpdateCall{
		service: s,
		iden:    iden,
		args:    make(map[string]interface{}),
	}
}

func (c *ChatUpdateCall) Muted(muted bool) *ChatUpdateCall {
	c.args["muted"] = muted
	return c
}

func (c *ChatUpdateCall) Do() (*Chat, error) {
	data, err := c.service.client.run("POST", "chats/"+c.iden, c.args)
	if err != nil {
		return nil, err
	}

	var chat Chat
	err = json.Unmarshal(data, &chat)
	if err != nil {
		return nil, err
	}
	return &chat, nil
}

type ChatDeleteCall struct {
	service *ChatService
	iden    string
}

func (s *ChatService) Delete(iden string) *ChatDeleteCall {
	return &ChatDeleteCall{
		service: s,
		iden:    iden,
	}
}

func (c *ChatDeleteCall) Do() error {
	_, err := c.service.client.run("DELETE", "chats/"+c.iden, nil)
	if err != nil {
		return err
	}
	return nil
}
