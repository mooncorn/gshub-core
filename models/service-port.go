package models

import (
	"time"

	"gorm.io/gorm"
)

type ServicePort struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deletedAt,omitempty"`
	ServiceID     uint           `gorm:"not null" json:"-"`
	ContainerPort int            `gorm:"not null" json:"containerPort"` // Container host number
	HostPort      int            `gorm:"not null" json:"hostPort"`      // Host port number
	Description   string         `gorm:"not null" json:"description"`
	Protocol      string         `json:"protocol"`
}
