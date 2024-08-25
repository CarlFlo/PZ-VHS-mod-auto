package database

import "gorm.io/gorm"

type IGUIs struct {
	Model
	IGUI_perks_Carpentry    string
	IGUI_perks_Cooking      string
	IGUI_perks_Farming      string
	IGUI_perks_Doctor       string
	IGUI_perks_Electricity  string
	IGUI_perks_Metalworking string
	IGUI_perks_Aiming       string
	IGUI_perks_Reloading    string
	IGUI_perks_Fishing      string
	IGUI_perks_Trapping     string
	IGUI_perks_Foraging     string
	IGUI_perks_Tailoring    string
	IGUI_perks_Mechanics    string
	IGUI_perks_Lightfooted  string
	IGUI_perks_Nimble       string
	IGUI_perks_Sneaking     string
	IGUI_perks_Axe          string
	IGUI_perks_Blunt        string
	IGUI_perks_SmallBlade   string
}

func (IGUIs) TableName() string {
	return "igui"
}

func (id *IGUIs) AfterCreate(tx *gorm.DB) error {
	// Log in debug DB maybe
	return nil
}

// Saves the data to the database
func (id *IGUIs) Save() {
	DB.Save(&id)
}
