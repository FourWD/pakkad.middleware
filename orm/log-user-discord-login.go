package orm

import (
	"sync"
	"time"
)

type LogUserDiscordLogin struct {
	ID       string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	UserID   string `json:"user_id" query:"user_id" gorm:"type:varchar(20);index;"`
	Username string `gorm:"-"`

	LoginDate     *time.Time `json:"login_date" query:"login_date" gorm:"default:null;type:datetime;"`
	LogoutDate    *time.Time `json:"logout_date" query:"logout_date" gorm:"default:null;type:datetime;"`
	CurrentStatus string    `gorm:"-"`
}

type SafeMapUser struct {
	mutex sync.Mutex
	Data  map[string]LogUserDiscordLogin
}

func (sm *SafeMapUser) Add(key string, user LogUserDiscordLogin) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	sm.Data[key] = user
}

func (sm *SafeMapUser) Get(key string) (LogUserDiscordLogin, bool) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	user, ok := sm.Data[key]
	return user, ok
}
