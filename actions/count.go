package actions

import (
	"time"
	"aws-architect/abstract"
	"flag"
	"fmt"
	"os"
	"bufio"
	"aws-architect/helpers"
)

type Counter struct {
	abstract.Action
	InProgress	bool
	Success			bool
	Message			string
	Skip				int
	Path				string
}

func (c Counter) Init() (bool) {
	c.InProgress = false
	c.Success = false
	c.Message = ""
	return  true
}

func (c Counter) Reset() (bool) {
	c.InProgress = false
	c.Success = false
	c.Message = ""
	return  true
}

func (c *Counter) Execute(logChannel chan string) (bool) {
	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
		}
	}()
	c.InProgress=true
	logChannel <- fmt.Sprintf("Counting rows from file : %s", c.Path)
	var file *os.File
	var err error
	if file, err = os.Open(c.Path); err == nil {
		defer file.Close()
		logChannel <- fmt.Sprintf("File : %s found ...", c.Path)
		var numRows int = 0
		scanner := bufio.NewScanner(file)
		if c.Skip > 0 {
			logChannel <- fmt.Sprintf("Skipping : %d row(s) ...", c.Skip)
			for i := 0; i < c.Skip; i++ {
				scanner.Scan()
				scanner.Text()
			}
		}
		for scanner.Scan() {
			scanner.Text()
			numRows++
		}
		logChannel <- fmt.Sprintf("Number of Rows : %d, found in file %s", numRows, c.Path)
		c.Message = fmt.Sprintf("File %s contains %d rows.", c.Path, numRows)
		c.Success=true
		c.InProgress=false
		defer func() {
			// recover from panic caused by writing to a closed channel
			if r := recover(); r != nil {
			}
		}()
		return  true
	}
	logChannel <- err.Error()
	c.Message = fmt.Sprintf("Error in count file rows : %s", err.Error())
	c.Success=false
	c.InProgress=false
	return false
}

func (c *Counter) IsInProgress() (bool) {
	return  c.InProgress
}


func (c *Counter) GetExitCode() (int) {
	for true {
		if ! c.InProgress {
			break
		}
		time.Sleep(time.Second * 5)
	}
	if c.Success {
		return 0
	}
	return  1
}

func (c *Counter) GetCommand() (string) {
	return  c.Command
}

func (c *Counter) GetName() (string) {
	return  c.Name
}

func (c *Counter) GetUsage() (string) {
	return  c.Usage
}

func (c *Counter) AcquireValues() (bool) {
	flag.StringVar(&c.Path, "path", "", "Full qualified file path")
	flag.IntVar(&c.Skip, "skip", 0, "Number of row to skip count from file top")
	flag.Parse()
	return  c.Path != ""
}

func (c *Counter) GetLastMessage() (string) {
	return  c.Message
}

func InitCounter() {
	var parm1 abstract.Parameter = abstract.Parameter{
		Name: "file",
		Description: "Full qualified file path",
		Mandatory: true,
		HasValue: true,
		SampleValue: "file-path",
	}
	var parm2 abstract.Parameter = abstract.Parameter{
		Name: "skip",
		Description: "Number of row to skip count from file top",
		Mandatory: false,
		HasValue: true,
		SampleValue: "nun-rows",
	}
	var Parameters 	[]abstract.Parameter = make([]abstract.Parameter, 0)
	Parameters = append(Parameters, parm1)
	Parameters = append(Parameters, parm2)
	var  CounterAction *Counter = new (Counter)
	CounterAction.Parameters= Parameters
	CounterAction.Name = "Sample File Rows Counter"
	CounterAction.Command= "count"
	CounterAction.Description= "Count rows in a File"
	CounterAction.Usage=helpers.DefineUsage(CounterAction.Command, CounterAction.Description, CounterAction.Parameters)
	abstract.RegisterAction(CounterAction)
}
