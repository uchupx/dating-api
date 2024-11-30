package dto

type AuthRequest struct {
	GrantTypePassword
	BasicAuthRequest
}

type BasicAuthRequest struct {
	GrantType    string `json:"grant_type" validate:"required"`
	ClientId     string `json:"client_id" validate:"required"`
	ClientSecret string `json:"client_secret" validate:"required"`
}

type GrantTypePassword struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}

type GrantTypeRefreshToken struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenRequest struct {
	GrantTypeRefreshToken
	BasicAuthRequest
}

type SignUpRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	ClientKey string `json:"client_key"`
}
