package gopushbullet

import (
	"encoding/json"
	"time"
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
	args    map[string]interface{}
}

func NewPushService(client *Client) *PushService {
	return &PushService{client}
}

func (s *PushService) List() *PushListCall {
	call := &PushListCall{
		service: s,
		args:    make(map[string]interface{}),
	}
	return call
}

func (c *PushListCall) Active(b bool) *PushListCall {
	c.args["active"] = b
	return c
}

func (c *PushListCall) ModifiedAfter(t time.Time) *PushListCall {
	c.args["modified_after"] = t.Unix()
	return c
}

func (c *PushListCall) Limit(i int) *PushListCall {
	c.args["limit"] = i
	return c
}

func (c *PushListCall) Cursor(s string) *PushListCall {
	c.args["cursor"] = s
	return c
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

type PushCreateCall struct {
	service *PushService
	args    map[string]interface{}
}

func (s *PushService) CreateNote() *PushCreateCall {
	call := s.create("note")
	return call
}

func (s *PushService) create(t string) *PushCreateCall {
	call := &PushCreateCall{
		service: s,
		args:    make(map[string]interface{}),
	}
	call.args["type"] = t
	return call
}

func (c *PushCreateCall) Do() (*Push, error) {
	data, err := c.service.client.run("POST", "pushes", c.args)
	if err != nil {
		return nil, err
	}

	var p Push
	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (c *PushCreateCall) Target(target string) *PushCreateCall {
	c.args["target"] = target
	return c
}

func (c *PushCreateCall) Title(title string) *PushCreateCall {
	c.args["title"] = title
	return c
}

func (c *PushCreateCall) Body(body string) *PushCreateCall {
	c.args["body"] = body
	return c
}

func (s *PushService) CreateLink() *PushCreateCall {
	call := s.create("link")
	return call
}

// Should be able to be used from CreateLink
func (c *PushCreateCall) Url(url string) *PushCreateCall {
	c.args["url"] = url
	return c
}

// Should only be able to use Body call and the file ones. Not title or url
func (s *PushService) CreateFile() *PushCreateCall {
	call := s.create("file")
	return call
}

func (c *PushCreateCall) FileName(filename string) *PushCreateCall {
	c.args["file_name"] = filename
	return c
}

func (c *PushCreateCall) FileType(filetype string) *PushCreateCall {
	c.args["file_type"] = filetype
	return c
}

func (c *PushCreateCall) FileUrl(fileurl string) *PushCreateCall {
	c.args["file_url"] = fileurl
	return c
}

type PushUpdateCall struct {
	service *PushService
	iden    string
	args    map[string]interface{}
}

func (s *PushService) Update(iden string) *PushUpdateCall {
	call := &PushUpdateCall{
		service: s,
		iden:    iden,
		args:    make(map[string]interface{}),
	}
	return call
}

func (c *PushUpdateCall) Dismissed(dismissed bool) *PushUpdateCall {
	c.args["dismissed"] = dismissed
	return c
}

func (c *PushUpdateCall) Do() (*Push, error) {
	data, err := c.service.client.run("POST", "pushes/"+c.iden, c.args)
	if err != nil {
		return nil, err
	}

	var p Push
	err = json.Unmarshal(data, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

type PushDeleteCall struct {
	service *PushService
	iden    string
}

func (s *PushService) Delete(iden string) *PushDeleteCall {
	call := &PushDeleteCall{
		service: s,
		iden:    iden,
	}
	return call
}

func (c *PushDeleteCall) Do() error {
	_, err := c.service.client.run("DELETE", "pushes/"+c.iden, nil)
	if err != nil {
		return err
	}
	return nil
}

type PushDeleteAllCall struct {
	service *PushService
}

func (s *PushService) DeleteAll() *PushDeleteAllCall {
	call := &PushDeleteAllCall{
		service: s,
	}
	return call
}

func (c *PushDeleteAllCall) Do() error {
	_, err := c.service.client.run("DELETE", "pushes", nil)
	if err != nil {
		return err
	}
	return nil
}
