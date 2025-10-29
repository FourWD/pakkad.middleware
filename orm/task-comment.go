package orm

import (
	"github.com/FourWD/middleware/model"
)

type TaskComment struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(20);primary_key"`
	model.GormModel

	TaskID  string `json:"task_id" query:"task_id" gorm:"type:varchar(20)"`
	Comment string `json:"comment" query:"comment" gorm:"type:text"`
}
