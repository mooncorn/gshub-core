package models

import (
	"time"
)

type ServiceVolume struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	ServiceID   uint      `gorm:"not null" json:"-"`
	Host        string    `gorm:"not null" json:"host"`
	Destination string    `gorm:"not null" json:"destination"`
	Description string    `gorm:"not null" json:"description"`
}
