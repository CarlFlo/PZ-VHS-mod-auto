package src

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/projectZomboidVHS/src/utils"
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

	if !utils.IsSkillInFilter(code) {
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
		output += fmt.Sprintf("%s, ", utils.GetPZSkill(skill))
	}
	output = output[:len(output)-2]
	output += "] <TN>\"" // TN = TapeName

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
