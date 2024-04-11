package dto

type TokenRes struct {
	TokenType             *string `json:"token_type"`
	AccessToken           *string `json:"access_token"`
	ExpiresIn             *int64  `json:"expires_in"`
	RefreshToken          *string `json:"refresh_token"`
	RefreshTokenExpiresIn *int64  `json:"refresh_token_expires_in"`
}
