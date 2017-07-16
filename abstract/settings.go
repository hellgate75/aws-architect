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

func LoadSettings() Settings {
	var settings Settings = Settings{}
	if file, err := os.Open(util.GetCurrentPath() + "/.settings"); err == nil {
		if bytes, err := ioutil.ReadAll(file); err == nil {
			if err = yaml.Unmarshal(bytes, &settings); err == nil {
				return settings
			} else {
				logger.Error(err)
			}
		} else {
			logger.Error(err)
		}
	} else {
		logger.Error(err)
	}
	return settings
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
