package token

// RefeshTokenForm refresh token form
type RefeshTokenForm struct {
	RefreshToken string `json:"refresh_token,omitempty" example:"ABCDEF1234"`
}
