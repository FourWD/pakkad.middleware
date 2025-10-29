package orm

import (
	"github.com/FourWD/middleware/model"
)

type TaskAssign struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	TaskID string `json:"task_id" query:"task_id" gorm:"type:varchar(36)"`
	UserID string `json:"user_id" query:"user_id" gorm:"type:varchar(36)"`
}
