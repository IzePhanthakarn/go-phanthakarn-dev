package guest

// RegisterForm register form
type RegisterForm struct {
	Email         string `json:"email"`
	Prefix        string `json:"prefix"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	PhoneNumber   string `json:"phone_number"`
	Birthday      string `json:"birthday"`
	CitizenID     string `json:"citizen_id"`
	Password      string `json:"password"`
}

// CheckUserForm check user form
type CheckUserForm struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	CitizenID string `json:"citizen_id"`
}

// LoginForm loginn form
type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResetPasswordForm struct {
	CitizenID   string `json:"citizen_id"`
	Birthday    string `json:"birthday"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
