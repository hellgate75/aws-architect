package bootstrap

import (
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/hellgate75/aws-architect/actions"
)

var Settings abstract.Settings

func InitActions() {
	actions.InitCounter()
	actions.InitS3CreateBucket()
	actions.InitS3DeleteBucket()
	//abstract.SaveSettings(Settings)
	Settings = abstract.LoadSettings()
}
