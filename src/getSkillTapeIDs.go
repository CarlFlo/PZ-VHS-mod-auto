package src

import (
	"bufio"
	"os"
	"strings"

	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/projectZomboidVHS/src/utils"
)

func getTapesWithSkillsFromFile(filepath, origin string, tapes *[]VHS) {

	// We care about the 'itemDisplayName' ID
	file, err := os.Open(filepath)
	if err != nil {
		malm.Fatal("%v", err)
	}
	defer file.Close()

	tape := VHS{}
	tape.New(origin)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, "itemDisplayName") {
			// Save the old, if the tape had any skills
			if tape.KeepOrDiscard() {
				*tapes = append(*tapes, tape)
			}
			// Delete the old
			tape = VHS{}
			tape.New(origin)

			tape.ID = utils.ExtractStringBetweenSep(line, "\"", "\"")
			if len(tape.ID) == 0 {
				malm.Fatal("no ID could be extracted from: '%s'", line)
			}
			continue
		}

		// The line we care about are longer than the usual lines, therefore, we can filter out lines that we know are too short.
		if len(line) < 80 {
			continue
		}

		checkForSkill(line, &tape)
		checkForRecipe(line, &tape)
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		malm.Fatal("Could not read lines: '%v'", err)
	}

	// last check
	if tape.KeepOrDiscard() {
		//malm.Info("Saved: '%s'", tape.ID)
		*tapes = append(*tapes, tape)
	}

}

func checkForSkill(line string, vhs *VHS) {

	// Extract the codes
	matchLeft := strings.Split(line, "codes = \"")
	if len(matchLeft) != 2 {
		// Found nothing
		return
	}

	matchRight := strings.Split(matchLeft[1], "\" },")
	if len(matchRight) != 2 {
		malm.Fatal("Found nothing on second split when we should have: '%s'", line)
		return
	}

	// matchRight[0] now contains the codes we want. Separated by comma(s)
	matches := matchRight[0]
	for {
		parts := strings.SplitN(matches, ",", 2)
		if len(parts) == 1 {
			// No more splits can be made
			vhs.AddSkillString(parts[0])
			break
		}

		vhs.AddSkillString(parts[0])
		matches = parts[1]
	}

}

// Only expects on recepe per line
func checkForRecipe(line string, vhs *VHS) {

	key := "RCP="

	matchLeft := strings.Split(line, key)
	if len(matchLeft) < 2 {
		// No recipe
		return
	}

	matchRight := strings.Split(matchLeft[1], "\" },")
	if len(matchRight) != 2 {
		malm.Fatal("Found nothing on second split when we should have: '%s'", line)
		return
	}

	// Check for additional
	if strings.Contains(matchRight[0], key) || strings.Contains(matchRight[0], ",") {
		malm.Fatal("Could not parse recipe. Unexpected format: '%s'", line)
		return
	}

	// Contains the recipe
	if len(matchRight[0]) != 0 {
		vhs.AddRecipe(matchRight[0])
	}
}
