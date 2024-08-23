package main

import (
	"fmt"

	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/projectZomboidVHS/src"
	"github.com/CarlFlo/projectZomboidVHS/src/config"
)

func init() {
	config.Load()
	malm.SetLogBitmask(63)
	malm.SetLogVerboseBitmask(0)
}

func main() {

	src.ParseSkillTapes()
	localisations := src.GetAvailableLocalisations()

	for _, locale := range localisations {
		fmt.Printf("Working on: %s\n", locale)
		// Visit the correct folder inside: D:\SteamLibrary\steamapps\common\ProjectZomboid\media\lua\shared\Translate

		// Fetch the values (EN example)
		// D:\SteamLibrary\steamapps\common\ProjectZomboid\media\lua\shared\Translate\EN
		// Check in: IG_UI_EN.txt
		// - To through the file until all values in 'skillFilter' in VHSStruct.go
		// - - Trim the whitespace and extract the values between the 'quotes'
		// - - Save in temporary map

		// Check in: Recorded_Media_EN.txt for the official name of the VHS tape
		// - Trim whitespace. DE file looks different from the rest, but should not be a problem
		// - What to do if there is no entry for a language? Fallback to english? Use our own?
	}

}
