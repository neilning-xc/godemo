package model

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	Email     string `gorm:"type:varchar(100);unique_index"`
	Brithday  *time.Time
	Gender    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type UserInfo struct {
	Id        uint64 `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Gender    int8   `json:"gender"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
