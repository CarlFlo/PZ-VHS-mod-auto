package utils

import (
	"fmt"
	"reflect"

	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/projectZomboidVHS/src/database"
)

// All skills - needs to match database -> IGUIs.go -> IGUIs struct
var skillFilter = map[string]string{
	"CRP": "IGUI_perks_Carpentry",
	"COO": "IGUI_perks_Cooking",
	"FRM": "IGUI_perks_Farming",
	"DOC": "IGUI_perks_Doctor",
	"ELC": "IGUI_perks_Electricity",
	"MTL": "IGUI_perks_Metalworking",
	"AIM": "IGUI_perks_Aiming",
	"REL": "IGUI_perks_Reloading",
	"FIS": "IGUI_perks_Fishing",
	"TRA": "IGUI_perks_Trapping",
	"FOR": "IGUI_perks_Foraging",
	"TAI": "IGUI_perks_Tailoring",
	"MEC": "IGUI_perks_Mechanics",
	"LFT": "IGUI_perks_Lightfooted",
	"NIM": "IGUI_perks_Nimble",
	"SNE": "IGUI_perks_Sneaking",
	"BAA": "IGUI_perks_Axe",
	"BUA": "IGUI_perks_Blunt",
	"SBA": "IGUI_perks_SmallBlade",
}

/*
// Placeholder. Skills I do not know the code for.
// Add them later, the code will notice if there are new "codes" that we do not handle

"Spear":       "IGUI_perks_Spear",
"LongBlade":   "IGUI_perks_LongBlade",
"SmallBlunt":  "IGUI_perks_SmallBlunt",
"Sprinting":   "IGUI_perks_Sprinting",
"Agility":     "IGUI_perks_Agility",
"Accuracy":    "IGUI_perks_Accuracy",
"Guard":       "IGUI_perks_Guard",
"Maintenance": "IGUI_perks_Maintenance",
"Firearm":     "IGUI_perks_Firearm",
"PPassiveAS":  "IGUI_perks_Passive",
"Strength":    "IGUI_perks_Strength",
"Fitness":     "IGUI_perks_Fitness",
"Combat":      "IGUI_perks_Combat",
"Survivalist": "IGUI_perks_Survivalist",
"Crafting":    "IGUI_perks_Crafting",
*/

var skillFilterExtra = map[string]string{
	"PAN": "IGUI_HaloNote_Panic",
	"UHP": "IGUI_HaloNote_Unhappiness",
	"STS": "IGUI_HaloNote_Stress",
}

// Returns the ID of the skill from the abbreviation
// CRP -> IGUI_perks_Carpentry
func GetPZSkill(abbreviation string) string {
	return skillFilter[abbreviation]
}

func IsSkillInFilter(skill string) bool {
	_, exists := skillFilter[skill]

	if _, e2 := skillFilterExtra[skill]; !e2 && !exists {
		malm.Info("Unknown VHS skill: '%s'", skill)
	}

	return exists
}

// the key will contain the skill ID, the value is intended to be the name in X language
func PopulateListWithSkills(list map[string]string) {

	for _, val := range skillFilter {
		list[val] = ""
	}
}

// Ensures the list of UGUI (skillFilter) matches the database struct IGUIs
func ValidateIGUI() error {

	structFields := make(map[string]struct{})

	// Use reflection to get the type of the IGUIs struct
	iguiType := reflect.TypeOf(database.IGUIs{})

	// Iterate over each field in the struct
	for i := 0; i < iguiType.NumField(); i++ {
		fieldName := iguiType.Field(i).Name
		structFields[fieldName] = struct{}{}
	}

	// Check if all values in the skillFilter map are in the struct fields
	for _, mapField := range skillFilter {
		if _, exists := structFields[mapField]; !exists {
			return fmt.Errorf("could not find map field '%s' in the database UGUI table. They need to match", mapField)
		}
	}

	// If all checks pass, return nil
	return nil
}
