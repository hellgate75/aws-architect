package helpers

import (
	"aws-architect/abstract"
	"fmt"
)

func DefineUsage(command string, description string, parameters []abstract.Parameter) (string) {
	var usage, args string
	usage=command + " "
	for i := 0; i < len(parameters); i++ {
		param := parameters[i];
		if param.Mandatory {
			usage += fmt.Sprintf("-%s %s ", param.Name, param.SampleValue)
			args += fmt.Sprintf("-%s %s\t\t\t%s\n", param.Name, param.SampleValue, param.Description)
		} else {
			usage += fmt.Sprintf("[-%s %s] ", param.Name, param.SampleValue)
			args += fmt.Sprintf("-%s %s\t\t\t%s [Optional]\n", param.Name, param.SampleValue, param.Description)
		}
	}
	usage += "\nDescription : " + description + "\nOptions:\n" + args
	return  usage
}