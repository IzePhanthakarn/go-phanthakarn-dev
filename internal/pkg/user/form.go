package user

import "github.com/IzePhanthakarn/go-phanthakarn-dev/internal/models"

type GetAllForm struct {
	models.PageForm
	Role *models.Role `json:"role" form:"role" query:"role"`
}
