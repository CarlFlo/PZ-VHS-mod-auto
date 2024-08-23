package src

import (
	"fmt"
	"os"
	"strings"

	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/projectZomboidVHS/src/config"
	"github.com/CarlFlo/projectZomboidVHS/src/utils"
)

func getLocalisations(folders *[]string) {

	// D:\SteamLibrary\steamapps\common\ProjectZomboid\media\lua\shared\Translate
	path := fmt.Sprintf("%s%s", config.CONFIG.GamePath, "\\media\\lua\\shared\\Translate")

	dir, err := os.Open(path)
	if err != nil {
		malm.Fatal("%v", err)
	}
	defer dir.Close()

	// Read the directory entries
	entries, err := dir.Readdir(0)
	if err != nil {
		malm.Fatal("%v", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {

			if !utils.Contains(config.CONFIG.IgnoreLocales, entry.Name()) {
				*folders = append(*folders, entry.Name())
			}
		}
	}

	// Move 'EN' to the first place in the list. We want to work on the English translation first because it is potentially a fallback
	enIndex := -1
	for i, folder := range *folders {
		if folder == "EN" {
			enIndex = i
			break
		}
	}
	if enIndex == -1 {
		malm.Fatal("Could not locate locale 'EN'")
	}
	// If "EN" is found and it's not already the first element, swap it with the first element
	if enIndex > 0 {
		(*folders)[0], (*folders)[enIndex] = (*folders)[enIndex], (*folders)[0]
	}
}

func GetAvailableLocalisations() []string {
	var folders []string
	getLocalisations(&folders)
	return folders
}

func ParseSkillTapes() {
	var tapes []VHS

	// D:\SteamLibrary\steamapps\common\ProjectZomboid\media\lua\shared\RecordedMedia
	filepath := fmt.Sprintf("%s\\%s", config.CONFIG.GamePath, "media\\lua\\shared\\RecordedMedia\\recorded_media.lua")
	getTapesWithSkillsFromFile(filepath, "Project Zomboid", &tapes)

	// Now we do all the mods
	for _, fPath := range config.CONFIG.ModRMPaths {
		fullPath := fmt.Sprintf("%s\\%s", config.CONFIG.WorkshopPath, fPath)
		origin := strings.Split(fPath, "\\")[1]

		getTapesWithSkillsFromFile(fullPath, origin, &tapes)
	}

	stringChannel := make(chan string)
	done := make(chan bool)

	go utils.WriteStringsFromChannel(stringChannel, "Recorded_Media_.txt", done)

	prevOrigin := ""
	for _, vhs := range tapes {
		if vhs.Origin != prevOrigin {
			if vhs.Origin != "Project Zomboid" {
				stringChannel <- ""
			}
			stringChannel <- fmt.Sprintf("// Tapes from %s", vhs.Origin)
		}
		stringChannel <- vhs.ToFormattedString()
		prevOrigin = vhs.Origin
	}

	close(stringChannel)
	if success := <-done; success {
		fmt.Println("Successfully wrote to the file.")
	} else {
		fmt.Println("An error occurred during file writing.")
	}

}
