package gopushbullet

import (
	"encoding/json"
)

type Upload struct {
	FileName  string `json:"file_name"`
	FileType  string `json:"file_type"`
	FileURL   string `json:"file_url"`
	UploadURL string `json:"upload_url"`
}

// POST https://api.pushbullet.com/v2/upload-request

type UploadService struct {
	client *Client
}

func NewUploadService(client *Client) *UploadService {
	return &UploadService{client}
}

type UploadUploadRequestCall struct {
	service *UploadService
	args    map[string]interface{}
}

func (s *UploadService) UploadRequest() *UploadUploadRequestCall {
	return &UploadUploadRequestCall{
		service: s,
		args:    make(map[string]interface{}),
	}
}

func (c *UploadUploadRequestCall) FileName(filename string) *UploadUploadRequestCall {
	c.args["file_name"] = filename
	return c
}

func (c *UploadUploadRequestCall) FileType(filetype string) *UploadUploadRequestCall {
	c.args["file_type"] = filetype
	return c
}

func (c *UploadUploadRequestCall) Do() (*Upload, error) {
	data, err := c.service.client.run("GET", "upload-request", c.args)
	if err != nil {
		return nil, err
	}

	var u Upload

	err = json.Unmarshal(data, &u)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
