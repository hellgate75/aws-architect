package helpers

import (
	"fmt"
	"github.com/hellgate75/aws-architect/abstract"
	"strings"
)

func calculateMaxParametersLength(parameters []abstract.Parameter) (int) {
	var length int = 0
	for i := 0; i < len(parameters); i++ {
		param := parameters[i]
		val := fmt.Sprintf("-%s %s", param.Name, param.SampleValue)
		if len(val) > length {
			length = len(val)
		}
	}
	return  length
}

func StringLPad(value string, pad int) (string) {
	var newValue string = value
	if pad > len(value) {
		newValue += strings.Repeat(" ", pad - len(value))
	}
	return  newValue
}


func DefineUsage(command string, description string, parameters []abstract.Parameter) string {
	var usage, args string
	var maxLen int = calculateMaxParametersLength(parameters)
	usage = command + " "
	for i := 0; i < len(parameters); i++ {
		param := parameters[i]
		var argument string = fmt.Sprintf("-%s %s", param.Name, param.SampleValue)
		var boundedArgument string = StringLPad(argument, maxLen)
		if param.Mandatory {
			usage += argument + " "
			args += fmt.Sprintf("%s    [Mandatory] %s\n", boundedArgument, param.Description)
		} else {
			usage += fmt.Sprintf("[%s] ", argument)
			args += fmt.Sprintf("%s    [Optional ] %s\n", boundedArgument, param.Description)
		}
	}
	usage += "\nDescription : " + description + "\nOptions:\n" + args
	return usage
}
