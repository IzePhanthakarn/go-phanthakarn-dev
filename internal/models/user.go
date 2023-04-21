package models

import "time"

// User this model is user
type User struct {
	Model
	Email         string      `json:"email"`
	Prefix        string      `json:"prefix"`
	FirstName     string      `json:"first_name"`
	LastName      string      `json:"last_name"`
	PhoneNumber   string      `json:"phone_number"`
	Birthday      time.Time   `json:"birthday"`
	CitizenID     string      `json:"citizen_id"`
	PasswordHash  string      `json:"-"`
	Roles         []*UserRole `json:"roles,omitempty"`
}

// RoleValues role values
func (u *User) RoleValues() []int {
	roles := []int{}
	for _, r := range u.Roles {
		roles = append(roles, int(r.Role))
	}
	return roles
}
