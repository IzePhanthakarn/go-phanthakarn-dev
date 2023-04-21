package models

// Role user role
type Role int

const (
	//CustomerRole customer [1]
	CustomerRole Role = iota + 1
)

const (
	//OfficerRole [10]
	OfficerRole Role = iota + 10
)

const (
	//AdminRole [100]
	AdminRole Role = iota + 100
)

const (
	//SuperAdminRole [1000]
	SuperAdminRole Role = iota + 1000
)

// UserRole user role
type UserRole struct {
	Model
	UserID uint `json:"-"`
	Role   Role `json:"role"`
}
