package util

import (
	"net/http"
  "io/ioutil"
	"github.com/hellgate75/aws-architect/log"
)
var logger log.Logger = log.Logger{}

func DownloadFile(filepath string, url string) (err error) {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		logger.WarningE(err)
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.WarningE(err)
		return  err
	}
	// Writer the body to file
	err = ioutil.WriteFile(filepath, body, 0666)
	if err != nil  {
		logger.WarningE(err)
		return err
	}

	return nil
}