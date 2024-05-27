package models

import (
	"time"

	"gorm.io/gorm"
)

type ImageEnvType struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	Name        string         `gorm:"not null" json:"name"`
	Key         string         `gorm:"uniqueIndex,not null" json:"key"`
	Required    bool           `gorm:"not null" json:"required"`
	Description string         `gorm:"not null" json:"description"`
	Default     string         `gorm:"not null" json:"default"`
	Category    string         `gorm:"not null" json:"category"`
}
