package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgress driver for gorm
	log "github.com/sirupsen/logrus"
)

// GetDb returns the database connection
func GetDb() *gorm.DB {
	dsn := fmt.Sprintf("sslmode=disable host=%s port=%d user=%s dbname=%s password=%s", cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbSchema, cfg.DbPassword)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Panicf("postgresql error: %v", err)
	}
	db.SingularTable(true)
	return db
}
