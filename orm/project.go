package orm

import (
	"time"

	"github.com/FourWD/middleware/model"
)

type Project struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(20);primary_key"` // Channel ID
	model.GormModel

	StartDate      *time.Time `json:"start_date" query:"start_date" gorm:"type:datetime"`
	Name           string     `json:"name" query:"name" gorm:"type:varchar(100)"`
	ProjectGroupID string     `json:"project_group_id" query:"project_group_id" gorm:"type:varchar(36)"`
}
