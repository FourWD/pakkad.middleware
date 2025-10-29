package orm

import (
	"time"
)

type LogUserDiscordJoinChannel struct {
	ID          string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	UserID      string `json:"user_id" query:"user_id" gorm:"type:varchar(20);index;"`
	Username    string `gorm:"-"`
	ChannelID   string `json:"channel_id" query:"channel_id" gorm:"type:varchar(20);index;"` // 299210103259922433
	ChannelName string `gorm:"-"`

	BeforeChannelID   string `gorm:"-"`
	BeforeChannelName string `gorm:"-"`

	LoginDate  time.Time `json:"login_date" query:"login_date" gorm:"default:null;type:datetime;"`
	LogoutDate time.Time `json:"logout_date" query:"logout_date" gorm:"default:null;type:datetime;"`
	IsLogout   bool      `json:"is_logout" query:"is_logout" gorm:"type:bool"`
}
