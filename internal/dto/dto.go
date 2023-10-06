package dto

import "github.com/iliyasali2107/wt-scheduler/internal/models"

type RegisterRequestBody struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Number     string `json:"number,omitempty"`
	TelegramId string `json:"telegram_id,omitempty"`
}

type RegisterResponseBody struct {
	User models.User
}

type LoginRequestBody struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginResponsBody struct {
	Token string `json:"token"`
}

type ValidateRequestBody struct {
	Token string `json:"token"`
}

type ValdiateResponseBody struct {
	UserId string `json:"user_id"`
}
