package models

type User struct {
	Id         int64
	Name       string `json:"name"`
	Password   string `json:"password"`
	Number     string `json:"number,omitempty"`
	TelegramId string `json:"telegram_id,omitempty"`
}
