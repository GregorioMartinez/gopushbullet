package gopushbullet

// Switch from panic to log
import (
	"encoding/json"
)

type Device struct {
	Active       bool    `json:"active"`
	AppVersion   float64 `json:"app_version"`
	Created      float64 `json:"created"`
	Iden         string  `json:"iden"`
	Manufacturer string  `json:"manufacturer"`
	Model        string  `json:"model"`
	Modified     float64 `json:"modified"`
	Nickname     string  `json:"nickname"`
	PushToken    string  `json:"push_token"`
}

type DeviceResponse struct {
	Devices []Device `json:"devices"`
}

type DeviceService struct {
	client *Client
}

type DeviceListCall struct {
	service *DeviceService
}

func NewDeviceService(client *Client) *DeviceService {
	return &DeviceService{client}
}

func (s *DeviceService) List() *DeviceListCall {
	return &DeviceListCall{
		service: s,
	}
}

func (c *DeviceListCall) Do() (*[]Device, error) {
	data, err := c.service.client.run("GET", "devices", nil)
	if err != nil {
		return nil, err
	}

	var d DeviceResponse
	err = json.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}

	return &d.Devices, nil
}
