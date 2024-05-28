package models

import (
	"time"

	"gorm.io/gorm"
)

type ServiceEnvValue struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	ServiceEnvID uint           `gorm:"not null" json:"serviceEnvId"`
	Value        string         `gorm:"not null" json:"value"`
	Name         string         `gorm:"not null" json:"name"`
}
