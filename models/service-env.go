package models

import (
	"time"

	"gorm.io/gorm"
)

type ServiceEnv struct {
	ID          uint              `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt    `gorm:"index" json:"deletedAt,omitempty"`
	ServiceID   uint              `json:"-"`
	Name        string            `gorm:"not null" json:"name"`
	Key         string            `gorm:"uniqueIndex,not null" json:"key"`
	Required    bool              `gorm:"not null" json:"required"`
	Description string            `gorm:"not null" json:"description"`
	Default     string            `gorm:"not null" json:"default"`
	Category    string            `gorm:"not null" json:"category"`
	Values      []ServiceEnvValue `gorm:"not null" json:"values"`
}