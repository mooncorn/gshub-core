package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GormDB struct {
	DB *gorm.DB
}

func NewGormDB(dsn string, config *gorm.Config) *GormDB {
	db, err := gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return &GormDB{DB: db}
}

func (g *GormDB) GetDB() *gorm.DB {
	return g.DB
}
