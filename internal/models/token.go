package models

// Token token model
type Token struct {
	Model
	Token        string `json:"token" gorm:"-"`
	RefreshToken string `json:"refresh_token"`
	DeviceToken  string `json:"device_token,omitempty"`
	UserID       uint   `json:"-"`
	User         *User  `json:"-"`
}
