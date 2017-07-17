package helpers

import (
	"os"
	"io/ioutil"
	"github.com/hellgate75/aws-architect/model"
	"github.com/hellgate75/aws-architect/log"
	"gopkg.in/yaml.v2"
)

var logger log.Logger = log.Logger{}

func LoadCORs(name string) (model.CORS, error) {
	var CORs model.CORS = model.CORS{}
	var err error
	var file *os.File
	if _,err = os.Stat(name); err == nil {
		if file, err = os.Open(name); err == nil {
			var bytes []byte = make([]byte, 0)
			if bytes, err = ioutil.ReadAll(file); err == nil {
				err = DeserializeCORS(bytes, &CORs)
				if err != nil {
					logger.WarningE(err)
				}
			} else {
				logger.WarningE(err)
			}
		} else {
			logger.WarningE(err)
		}
	} else {
		logger.WarningE(err)
	}
	return CORs, err
}
 func DeserializeCORS(bytes []byte, CORs *model.CORS) (error) {
	 var err error
	 if err = yaml.Unmarshal(bytes, CORs); err == nil {
		 return nil
	 } else {
		 logger.Error(err)
		 return err
	 }
 }
