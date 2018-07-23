package models

import "github.com/raj47i/IntTestPlivo/config"

// PhoneNumber represents an HTTP BasicAuth account
type PhoneNumber struct {
	ID        uint `gorm:"primary_key"`
	Number    string
	AccountID uint
	Account   Account `gorm:"foreignkey:AccountID"`
}

// LoadByNumberAndAccountID populates object from the DB based on number and account_id
func (pn *PhoneNumber) LoadByNumberAndAccountID(number string, aid uint) (bool, error) {
	db := config.GetDb()
	defer db.Close()
	if r := db.Where("number = ? AND account_id =?", number, aid).First(&pn); r.Error != nil {
		if r.RecordNotFound() {
			return false, nil
		}
		return false, r.Error
	}
	return pn.ID > 0, nil
}
