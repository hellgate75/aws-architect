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
	actions.InitS3BucketStatus()
	actions.InitListS3Buckets()
	//abstract.SaveSettings(Settings)

	var err error
	Settings, err = abstract.LoadSettings()
	if err != nil {
		InitDatabaseConfig(&Settings)
	}
}
