package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionId string    `gorm:"type:char(36);primary_key" json:"transaction_id"`
	User          User      `gorm:"foreignKey:UserId;references:UserId" json:"user"`
	UserId        uuid.UUID `gorm:"type:char(36);primary_key" json:"user_id"`
	Total         float64   `json:"total"`
	CreatedAt     time.Time `gorm:"type:timestamp;not null" json:"created_at"`
	
}
