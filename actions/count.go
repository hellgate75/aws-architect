package actions

import (
	"github.com/hellgate75/aws-architect/abstract"
	"github.com/hellgate75/aws-architect/command"
	"github.com/hellgate75/aws-architect/helpers"
)

func InitCounter() {
	var parm1 abstract.Parameter = abstract.Parameter{
		Name:        "file",
		Description: "Full qualified file path",
		Mandatory:   true,
		HasValue:    true,
		SampleValue: "file-path",
	}
	var parm2 abstract.Parameter = abstract.Parameter{
		Name:        "skip",
		Description: "Number of row to skip count from file top",
		Mandatory:   false,
		HasValue:    true,
		SampleValue: "nun-rows",
	}
	var Parameters []abstract.Parameter = make([]abstract.Parameter, 0)
	Parameters = append(Parameters, parm1)
	Parameters = append(Parameters, parm2)
	var CounterAction *abstract.ActionImpl = new(abstract.ActionImpl)
	CounterAction.Parameters = Parameters
	CounterAction.Name = "Sample File Rows Counter"
	CounterAction.Command = "count"
	CounterAction.Description = "Count rows in a File"
	CounterAction.Usage = helpers.DefineUsage(CounterAction.Command, CounterAction.Description, CounterAction.Parameters)
	CounterAction.SetArgumentParser(new(command.CounterParser))
	CounterAction.SetExecutor(new(command.CounterCommand))
	abstract.RegisterAction(CounterAction)
}
