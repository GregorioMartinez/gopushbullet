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

func NewDeviceService(client *Client) *DeviceService {
	return &DeviceService{client}
}

type DeviceListCall struct {
	service *DeviceService
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

type DeviceCreateCall struct {
	service *DeviceService
	args    map[string]interface{}
}

// Required go in as params.
// Optional are additional methods
//@TODO - Contact them about this?
func (s *DeviceService) Create() *DeviceCreateCall {

	call := &DeviceCreateCall{
		service: s,
		args:    make(map[string]interface{}),
	}

	return call
}

type DeviceUpdateCall struct {
	service *DeviceService
	iden    string
	args    map[string]interface{}
}

func (s *DeviceService) Update(iden string) *DeviceUpdateCall {
	call := &DeviceUpdateCall{
		service: s,
		iden:    iden,
		args:    make(map[string]interface{}),
	}
	return call
}

func (c *DeviceUpdateCall) Nickname(name string) *DeviceUpdateCall {
	c.args["nickname"] = name
	return c
}

func (c *DeviceUpdateCall) Model(model string) *DeviceUpdateCall {
	c.args["model"] = model
	return c
}

func (c *DeviceUpdateCall) Manufacturer(manufacturer string) *DeviceUpdateCall {
	c.args["manufacturer"] = manufacturer
	return c
}

func (c *DeviceUpdateCall) PushToken(push_token string) *DeviceUpdateCall {
	c.args["push_token"] = push_token
	return c
}

func (c *DeviceUpdateCall) AppVersion(app_version int) *DeviceUpdateCall {
	c.args["app_version"] = app_version
	return c
}

func (c *DeviceUpdateCall) Icon(icon string) *DeviceUpdateCall {
	c.args["icon"] = icon
	return c
}

func (c *DeviceUpdateCall) HasSMS(has_sms bool) *DeviceUpdateCall {
	c.args["has_sms"] = has_sms
	return c
}

func (c *DeviceUpdateCall) Do() (*Device, error) {

	data, err := c.service.client.run("POST", "devices/"+c.iden, c.args)
	if err != nil {
		return nil, err
	}

	var d Device
	err = json.Unmarshal(data, &d)
	if err != nil {
		return nil, err
	}

	return &d, nil
}
