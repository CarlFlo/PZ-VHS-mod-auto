package src

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CarlFlo/malm"
)

type VHS struct {
	ID      string
	Skills  map[string]int
	Recipes []string
	Origin  string
}

func (t *VHS) New(origin string) {
	t.Skills = make(map[string]int)
	t.Origin = origin
}

// Expected input format: 'COO+1'
func (t *VHS) AddSkillString(skill string) {

	parts := strings.Split(skill, "+")
	if len(parts) == 1 {
		return
	}
	code := parts[0]

	if !isSkillInFilter(code) {
		return
	}

	number, err := strconv.Atoi(parts[1])
	if err != nil {
		malm.Fatal("Could not parse number from '%s'", parts[1])
	}

	t.AddSkill(code, number)
}

func (t *VHS) AddSkill(skill string, val int) {
	t.Skills[skill] += val
}

func (t *VHS) AddRecipe(recipe string) {
	t.Recipes = append(t.Recipes, recipe)
}

func (t *VHS) KeepOrDiscard() bool {
	return len(t.Skills) > 0
}

func (t *VHS) ToFormattedString() string {

	output := fmt.Sprintf("%s = \"[", t.ID)

	for skill := range t.Skills {
		output += fmt.Sprintf("%s, ", skillFilter[skill])
	}
	output = output[:len(output)-2]
	output += "] <TAPENAME>\""

	return output
}

func (t *VHS) ToString() string {
	out := fmt.Sprintf("[%s] %s - ", t.Origin, t.ID)
	for skill, value := range t.Skills {
		out += fmt.Sprintf("%s (%d), ", skill, value)
	}
	out = out[:len(out)-2]

	if len(t.Recipes) != 0 {
		out = fmt.Sprintf("%s - ", out)
		for _, recipe := range t.Recipes {
			out += fmt.Sprintf("%s, ", recipe)
		}
		out = out[:len(out)-2]
	}

	return out
}

// All skills
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

func isSkillInFilter(skill string) bool {
	_, exists := skillFilter[skill]

	if _, e2 := skillFilterExtra[skill]; !e2 && !exists {
		malm.Info("Unknown VHS skill: '%s'", skill)
	}

	return exists
}
