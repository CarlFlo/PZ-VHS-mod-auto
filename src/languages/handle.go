package languages

import (
	"bufio"
	"fmt"
	"os"
	"regexp"

	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/projectZomboidVHS/src/config"
	"github.com/CarlFlo/projectZomboidVHS/src/utils"
)

func ParseLanguageData(localisations *[]string) {

	for _, locale := range *localisations {
		fmt.Printf("Working on: %s\n", locale)

		langTemp := make(map[string]string)

		handle(locale, langTemp)
	}

}

func handle(language string, lang map[string]string) {

	IGUIpath := fmt.Sprintf("%s\\media\\lua\\shared\\Translate\\%s\\IG_UI_%s.txt", config.CONFIG.GamePath, language, language)
	RMpath := fmt.Sprintf("%s\\media\\lua\\shared\\Translate\\%s\\Recorded_Media_%s.txt", config.CONFIG.GamePath, language, language)

	handleIGUI(IGUIpath)
	handleRM(RMpath)
}

type checkList struct {
	list map[string]string
	left int
}

func (cl *checkList) New() {
	cl.list = make(map[string]string)
	utils.PopulateListWithSkills(cl.list)
	cl.left = len(cl.list)
}

func (cl *checkList) CheckLine(line string) bool {

	// We only care about 'IGUI_perks_'
	r := regexp.MustCompile(`\s*(IGUI_perks_\w+)\s*=\s*"(.*?)"`)

	// Find all matches
	matches := r.FindStringSubmatch(line)

	if len(matches) != 3 {
		return false
	}

	perkID := matches[1]
	value := matches[2]

	fmt.Printf("%s = %s\n", perkID, value)
	return cl.Add(perkID, value) // Add logic here for check
}

func (cl *checkList) Add(perkID, value string) bool {
	return false
}

func handleIGUI(filepath string) {
	// Fetch the values (EN example)
	// D:\SteamLibrary\steamapps\common\ProjectZomboid\media\lua\shared\Translate\EN
	// Check in: IG_UI_EN.txt
	// - To through the file until all values in 'skillFilter' in VHSStruct.go
	// - - Trim the whitespace and extract the values between the 'quotes'
	// - - Save in temporary map

	var cl checkList
	cl.New()

	file, err := os.Open(filepath)
	if err != nil {
		malm.Fatal("%v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		done := cl.CheckLine(line)
		if done {
			break
		}
	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		malm.Fatal("Could not read lines: '%v'", err)
	}

}

func handleRM(path string) {
	// Check in: Recorded_Media_EN.txt for the official name of the VHS tape
	// - Trim whitespace. DE file looks different from the rest, but should not be a problem
	// - What to do if there is no entry for a language? Fallback to english? Use our own?

}
