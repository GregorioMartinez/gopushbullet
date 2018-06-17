package gopushbullet

type ErrorResp struct {
	Error Error
}

type Error struct {
	Cat     string `json:"cat"`
	Message string `json:"message"`
	Type    string `json:"type"`
}

func (e Error) Error() string {
	return e.Message
}
