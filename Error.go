package gopushbullet

type Error struct {
	Cat     string `json:"cat"`
	Message string `json:"message"`
	Type    string `json:"type"`
}
