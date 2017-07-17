package abstract

import (
	"time"
)

type ActionImpl struct {
	Action
	InProgress bool
	Success    bool
	Message    string
	Parser		 ArgumentParser
	Executor		 Command
}

func (c ActionImpl) Init() bool {
	c.InProgress = false
	c.Success = false
	c.Message = ""
	return true
}

func (c ActionImpl) Reset() bool {
	c.InProgress = false
	c.Success = false
	c.Message = ""
	return true
}

func (c *ActionImpl) Execute(logChannel chan string) bool {
	return c.Executor.Execute(c, c.Parser.Parse(), logChannel)
}

func (c *ActionImpl) IsInProgress() bool {
	return c.InProgress
}

func (c *ActionImpl) GetExitCode() int {
	for true {
		if !c.InProgress {
			break
		}
		time.Sleep(time.Second * 5)
	}
	if c.Success {
		return 0
	}
	return 1
}

func (c *ActionImpl) GetCommand() string {
	return c.Command
}

func (c *ActionImpl) GetName() string {
	return c.Name
}

func (c *ActionImpl) GetUsage() string {
	return c.Usage
}

func (c *ActionImpl) SetArgumentParser(parser ArgumentParser) () {
	c.Parser = parser
}

func (c *ActionImpl) SetExecutor(command Command) () {
	c.Executor = command
}

func (c *ActionImpl) AcquireValues() bool {
	return c.Parser.Validate()
}

func (c *ActionImpl) GetLastMessage() string {
	return c.Message
}

