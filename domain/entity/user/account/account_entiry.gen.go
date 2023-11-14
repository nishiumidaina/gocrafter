package account

import (
	"time"
)

type Accounts []Account

type Account struct {
	ID int64 `json:"id"`

	UUID int64 `json:"uuid"`

	Name string `json:"name"`

	Password string `json:"password"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`

	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
