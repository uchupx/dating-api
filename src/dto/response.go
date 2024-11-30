package dto

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta"`
	Data    interface{} `json:"data"`
}

type TokenResponse struct {
	Token        string  `json:"access_token"`
	TokenType    string  `json:"token_type"`
	Expired      int64   `json:"expired"`
	RefreshToken string  `json:"refresh_token"`
	Scope        *string `json:"scope"`
}

type EntityResponse struct {
	Id     interface{} `json:"id"`
	Entity string      `json:"entity"`
	Meta   interface{} `json:"meta"`
}
