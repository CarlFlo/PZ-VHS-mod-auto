package main

import (
	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/projectZomboidVHS/src"
	"github.com/CarlFlo/projectZomboidVHS/src/config"
	"github.com/CarlFlo/projectZomboidVHS/src/database"
	"github.com/CarlFlo/projectZomboidVHS/src/languages"
	"github.com/CarlFlo/projectZomboidVHS/src/utils"
)

const CurrentVersion = "2024-08-23"

func init() {
	malm.SetLogBitmask(63)
	malm.SetLogVerboseBitmask(0)

	config.Load()
	database.Connect()
	go utils.CheckVersion(CurrentVersion)
}

func main() {
	src.ParseSkillTapes()
	localisations := src.GetAvailableLocalisations()

	languages.ParseLanguageData(&localisations)
}
