package dto

type RegisterRequestBody struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	Number     string `json:"number,omitempty"`
	TelegramId string `json:"telegram_id,omitempty"`
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
