package database

import "gorm.io/gorm"

type Lang struct {
	Model
}

func (Lang) TableName() string {
	return "lang"
}

func (l *Lang) AfterCreate(tx *gorm.DB) error {
	// Log in debug DB maybe
	return nil
}

// Saves the data to the database
func (l *Lang) Save() {
	DB.Save(&l)
}
