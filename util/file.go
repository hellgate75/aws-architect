package util

import (
	"os"
)

func GetCurrentPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return pwd
}
