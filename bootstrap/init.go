package bootstrap

import (
	"aws-architect/actions"
	"aws-architect/abstract"
)

var Settings abstract.Settings

func InitActions() {
	actions.InitCounter()
	//abstract.SaveSettings(Settings)
	Settings = abstract.LoadSettings()
}
