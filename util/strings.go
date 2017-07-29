package util

import "strings"

func StringRPad(value string, pad int) (string) {
	var newValue string = value
	if pad > len(value) {
		newValue += strings.Repeat(" ", pad - len(value))
	}
	return  newValue
}


func StringLPad(value string, pad int) (string) {
	var newValue string = value
	if pad > len(value) {
		newValue = strings.Repeat(" ", pad - len(value)) + newValue
	}
	return  newValue
}
