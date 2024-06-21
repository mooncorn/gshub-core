package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type GormDB struct {
	DB *gorm.DB
}

func NewPostgresDB(dsn string, config *gorm.Config) *GormDB {
	return NewGormDB(postgres.Open(dsn), config)
}

func NewSQLiteDB(dbFile string, config *gorm.Config) *GormDB {
	return NewGormDB(sqlite.Open(dbFile), config)
}

func NewGormDB(dialector gorm.Dialector, config *gorm.Config) *GormDB {
	db, err := gorm.Open(dialector, config)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return &GormDB{DB: db}
}

func (g *GormDB) GetDB() *gorm.DB {
	return g.DB
}
