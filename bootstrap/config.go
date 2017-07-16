package bootstrap

import (
	"os"
	"github.com/blakesmith/ar"
	"bytes"
	"io"
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/hellgate75/aws-architect/util"
	"github.com/hellgate75/aws-architect/log"
	"fmt"
	"io/ioutil"
)
var logger log.Logger = log.Logger{}

func InitDatabaseConfig(settings *abstract.Settings) {
	var arFileName string = util.GetCurrentPath() + "/aws-db.a"
	if _,err := os.Stat(arFileName); err == nil {
		err, xMap := ReadConfig(arFileName)
		if err == nil {
			if val, ok := xMap[".settings"]; ok {
				abstract.DeserializeSettings(val, settings)
				logger.Log(fmt.Sprintf("Configuration reloaded from file : %s", arFileName))
				ioutil.WriteFile(util.GetCurrentPath() + "/.settings", val, 0666)
			}
		}
	} else  {
		LoadFromURL(arFileName, settings)
	}
}

func ReadConfig(name string) (error, map[string][]byte) {
	var files map[string][]byte = make(map[string][]byte, 0)
	file, err := os.Open(name)
	if err != nil {
		return  err, files
	}
	reader := ar.NewReader(file)

	if header, errH := reader.Next(); errH==nil {
		var buf bytes.Buffer
		io.Copy(&buf, reader)
		files[header.Name] = buf.Bytes()
	}
	return nil, files
}

func LoadFromURL(arFileName string, settings *abstract.Settings) {
	var url string = "https://raw.githubusercontent.com/hellgate75/aws-architect/master/aws-db.a"
	if _,err := os.Stat(arFileName); err !=nil {
		util.DownloadFile(arFileName, url)
	}
	err,xMap := ReadConfig(arFileName)
	if err == nil {
		logger.Log(fmt.Sprintf("Database loaded from url: %s", url))
		if val, ok := xMap[".settings"]; ok {
			abstract.DeserializeSettings(val, settings)
			logger.Log(fmt.Sprintf("Configuration reloaded from file : %s", arFileName))
			ioutil.WriteFile(util.GetCurrentPath() + "/.settings", val, 0666)
		}
	}
}
