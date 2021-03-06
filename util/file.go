package util

import (
	"io/ioutil"
	"os"
)

func GetCurrentPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return pwd
}

func LoadFileContent(path string) (string, error) {
	var err error
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(bytes), err
}
