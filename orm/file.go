package orm

import (
	"github.com/FourWD/middleware/model"
)

type File struct { 
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	Name string `json:"name" query:"name" gorm:"type:varchar(250)"`
	Path string `json:"path" query:"path" gorm:"type:varchar(500)"`
}
