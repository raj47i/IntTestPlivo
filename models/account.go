package models

import (
	"github.com/raj47i/IntTestPlivo/config"
)

// Account represents an HTTP BasicAuth account
type Account struct {
	ID       uint `gorm:"primary_key"`
	AuthID   string
	Username string
}

// LoadByUserName populates object from the DB based on username
func (acc *Account) LoadByUserName(uname string) bool {
	db := config.GetDb()
	defer db.Close()
	db.Where("username = ?", uname).First(&acc)
	return acc.ID > 0
}

// Authenticate validates user's auth_id (secret)
func (acc *Account) Authenticate(secret string) bool {
	return acc.Username != "" && acc.AuthID == secret
}
