package bootstrap

import (
	"aws-architect/abstract"
	"aws-architect/actions"
)

var Settings abstract.Settings

func InitActions() {
	actions.InitCounter()
	actions.InitS3CreateBucket()
	actions.InitS3DeleteBucket()
	//abstract.SaveSettings(Settings)
	Settings = abstract.LoadSettings()
}
