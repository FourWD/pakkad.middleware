package orm

import (
	"github.com/FourWD/middleware/model"
)

type User struct {
	ID            string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	DiscordUserID string `json:"discord_user_id" query:"discord_user_id" gorm:"type:varchar(20);unique;"`
	model.GormModel

	UserTypeID string `json:"user_type_id" query:"user_type_id" gorm:"type:varchar(2)"`
	Username   string `json:"username" query:"username" gorm:"type:varchar(10)"`
	Password   string `json:"password" query:"password" gorm:"type:varchar(100)"`
	Firstname  string `json:"firstname" query:"firstname" gorm:"type:varchar(50)"`
	Lastname   string `json:"lastname" query:"lastname" gorm:"type:varchar(50)"`
}
