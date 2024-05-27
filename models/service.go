package models

import (
	"time"

	"gorm.io/gorm"
)

type ImageVariableConfigurationType string

const (
	MemoryConfigurationType ImageVariableConfigurationType = "MEMORY"
	ServerConfigurationType ImageVariableConfigurationType = "SERVER"
)

type ImageVariable struct {
	ID                uint                           `gorm:"primaryKey" json:"id"`
	CreatedAt         time.Time                      `json:"createdAt"`
	UpdatedAt         time.Time                      `json:"updatedAt"`
	DeletedAt         gorm.DeletedAt                 `gorm:"index" json:"deletedAt,omitempty"`
	Name              string                         `gorm:"uniqueIndex,not null" json:"name"` // Unique identifier
	Required          bool                           `gorm:"not null"`
	Description       string                         `gorm:"not null"`
	Default           string                         `gorm:"not null"`
	Category          string                         `gorm:"not null"`
	ConfigurationType ImageVariableConfigurationType `gorm:"not null"`
}
