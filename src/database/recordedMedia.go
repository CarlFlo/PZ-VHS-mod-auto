package database

import "gorm.io/gorm"

type RecordedMedia struct {
	Model
	RecodedMediaID string
}

func (RecordedMedia) TableName() string {
	return "recordedMedia"
}

func (rm *RecordedMedia) AfterCreate(tx *gorm.DB) error {
	// Log in debug DB maybe
	return nil
}

// Saves the data to the database
func (rm *RecordedMedia) Save() {
	DB.Save(&rm)
}
