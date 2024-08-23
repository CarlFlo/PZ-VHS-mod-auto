package config

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/CarlFlo/malm"
)

// CONFIG holds all the config data
var CONFIG *configStruct

type configStruct struct {
	GamePath      string   `json:"gamePath"`
	WorkshopPath  string   `json:"workshopPath"`
	IgnoreLocales []string `json:"ignoreLocales"`
	ModRMPaths    []string `json:"modRMPaths"`
	Database      database `json:"database"`
}

type database struct {
	FileName string `json:"fileName"`
}

// ReloadConfig is a wrapper function for reloading the config. For clarity
func ReloadConfig() error {
	return readConfig()
}

// readConfig will read the config file
func readConfig() error {

	file, err := os.Open("./config.json")
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	file.Close()

	if err = json.Unmarshal(buf.Bytes(), &CONFIG); err != nil {
		return err
	}

	return nil
}

// createConfig creates the default config file
func createConfig() error {

	// Default config settings
	configStruct := configStruct{
		GamePath:      "D:\\SteamLibrary\\steamapps\\common\\ProjectZomboid",
		WorkshopPath:  "D:\\SteamLibrary\\steamapps\\workshop\\content",
		IgnoreLocales: []string{"RO"},
		ModRMPaths: []string{
			"108600\\3153010942\\mods\\FirstAidVHSTapes\\media\\lua\\shared\\RecordedMedia\\recorded_media_FirstAidVHS.lua",
			"108600\\2702055974\\mods\\SkillTapes\\media\\lua\\shared\\RecordedMedia\\SkillTapes_recorded_media.lua"},
		Database: database{
			FileName: "database.db",
		},
	}

	jsonData, _ := json.MarshalIndent(configStruct, "", "   ")
	err := os.WriteFile("config.json", jsonData, 0644)

	return err
}

// loadConfiguration loads the configuration file into memory
func loadConfiguration() error {

	if err := readConfig(); err != nil {
		if err = createConfig(); err != nil {
			return err
		}
		if err = readConfig(); err != nil {
			return err
		}
	}
	return nil
}

func Load() {
	if err := loadConfiguration(); err != nil {
		malm.Fatal("Error loading configuration: %s", err)
		return
	}

	requiredVariableCheck()

	//malm.Info("Configuration loaded")
}

// Some variables are required for the bot to work
func requiredVariableCheck() {

	// This function checks if some important variables are set in the config file
	problem := false

	if len(CONFIG.GamePath) == 0 {
		malm.Error("No path to project zomboid provided!")
		problem = true
	}

	if problem {
		malm.Fatal("There are at least one variable missing in the configuration file. Please fix the above errors!")
	}
}
