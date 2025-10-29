package orm

import (
	"github.com/FourWD/middleware/model"
)

type LogActivity struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	ProjectID string `json:"project_id" query:"project_id" gorm:"type:varchar(36)"`
	TaskID    string `json:"task_id" query:"task_id" gorm:"type:varchar(36)"`
	Remark    string `json:"remark" query:"remark" gorm:"type:varchar(500)"`
}
