package abstract

import (
	"github.com/hellgate75/aws-architect/log"
	"github.com/hellgate75/aws-architect/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var logger log.Logger = log.Logger{}

type Settings struct {
	DebugDisabled bool `yaml:"noDebug"`
}

func LoadSettings()  (Settings, error) {
	var settings Settings = Settings{}
	var err error
	var file *os.File
	if file, err = os.Open(util.GetCurrentPath() + "/.settings"); err == nil {
		var bytes []byte = make([]byte, 0)
		if bytes, err = ioutil.ReadAll(file); err == nil {
			err = DeserializeSettings(bytes, &settings)
			if err != nil {
				logger.WarningE(err)
			}
		} else {
			logger.WarningE(err)
		}
	} else {
		logger.WarningE(err)
	}
	return settings, err
}

func DeserializeSettings(bytes []byte, settings *Settings) (error) {
	var err error
	if err = yaml.Unmarshal(bytes, settings); err == nil {
		return nil
	} else {
		logger.Error(err)
		return err
	}
}

func SaveSettings(settings Settings) error {
	settings.DebugDisabled = true
	if bytes, err := yaml.Marshal(settings); err == nil {
		if err := ioutil.WriteFile(util.GetCurrentPath()+"/.settings", bytes, 777); err == nil {
			println("File written")
		} else {
			logger.Error(err)
			return err
		}
	} else {
		logger.Error(err)
		return err
	}
	return nil
}
