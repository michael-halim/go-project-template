package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserId       uuid.UUID `gorm:"type:char(36);primary_key" json:"user_id"`
	Username     string    `gorm:"unique,type:varchar(255);not null" json:"username"`
	Email        string    `gorm:"unique,type:varchar(255);not null" json:"email"`
	Password     string    `gorm:"type:varchar(255);not null" json:"-"`
	Salt         string    `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt    time.Time `gorm:"type:timestamp;not null" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt    time.Time `gorm:"type:timestamp" json:"deleted_at"`
	LastLogin    time.Time `gorm:"type:timestamp" json:"last_login"`
	IsSuperadmin bool      `gorm:"type:boolean;not null" json:"is_superadmin"`
	Status       uint8     `json:"status"`
}
