package main

import (
	"github.com/CarlFlo/malm"
	"github.com/CarlFlo/projectZomboidVHS/src"
	"github.com/CarlFlo/projectZomboidVHS/src/config"
	"github.com/CarlFlo/projectZomboidVHS/src/languages"
)

func init() {
	config.Load()
	malm.SetLogBitmask(63)
	malm.SetLogVerboseBitmask(0)
}

func main() {
	src.ParseSkillTapes()
	localisations := src.GetAvailableLocalisations()

	languages.ParseLanguageData(&localisations)
}
