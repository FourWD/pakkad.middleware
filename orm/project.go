package orm

import (
	"github.com/FourWD/middleware/model"
)

type Project struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(20);primary_key"` // Channel ID
	model.GormModel

	Name string `json:"name" query:"name" gorm:"type:varchar(100)"`
}
