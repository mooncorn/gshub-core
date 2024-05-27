package models

import (
	"time"

	"gorm.io/gorm"
)

type ImageEnvValue struct {
	ID             uint           `gorm:"primaryKey" json:"id"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	ImageEnvTypeID uint           `json:"-"`
	Value          string         `json:"value"`
}
