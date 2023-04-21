package user

import "github.com/IzePhanthakarn/go-boilerplate/internal/models"

type GetAllForm struct {
	models.PageForm
	Role *models.Role `json:"role" form:"role" query:"role"`
}
