package bootstrap

import (
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/hellgate75/aws-architect/actions"
	"github.com/hellgate75/aws-architect/util"
	"github.com/golang/go/src/pkg/os"
)

var Settings abstract.Settings

func InitActions() {
	actions.InitCounter()
	actions.InitS3CreateBucket()
	actions.InitS3DeleteBucket()
	//abstract.SaveSettings(Settings)
	var err error
	Settings, err = abstract.LoadSettings()
	if err != nil {
		var arFileName string = util.GetCurrentPath() + "/aws-db.a"
		if _,err := os.Stat(arFileName); err !=nil {
			//TODO: Download from github.com
			// from URL : https://raw.githubusercontent.com/hellgate75/aws-architect/master/aws-architect.a
			util.DownloadFile(arFileName, "https://raw.githubusercontent.com/hellgate75/aws-architect/master/aws-db.a")
		}
		err,xMap := ReadConfig(arFileName)
		if err == nil {
			if val, ok := xMap[".settings"]; ok {
				abstract.DeserializeSettings(val, &Settings)
			}
		}
	}
}
