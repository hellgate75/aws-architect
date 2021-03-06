package bootstrap

import (
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/hellgate75/aws-architect/actions"
)

var Settings abstract.Settings

func InitActions() {
	//actions.InitCounter()
	actions.InitS3CreateBucket()
	actions.InitS3DeleteBucket()
	actions.InitS3BucketStatus()
	actions.InitListS3Buckets()
	actions.InitUploadToS3Bucket()
	//abstract.SaveSettings(Settings)

	actions.InitCreateCfnStack()

	var err error
	Settings, err = abstract.LoadSettings()
	if err != nil {
		InitDatabaseConfig(&Settings)
	}
}
