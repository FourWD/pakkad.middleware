package orm

import (
	"github.com/FourWD/middleware/model"
)

type ProjectOwner struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	ProjectID string `json:"project_id" query:"project_id" gorm:"type:varchar(36)"`
	UserID    string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
}
