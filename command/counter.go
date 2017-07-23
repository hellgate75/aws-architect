package command

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/hellgate75/aws-architect/abstract"
	"os"
	"strconv"
)

type CounterCommand struct {
}

func (p *CounterCommand) Execute(action *abstract.ActionImpl, arguments []interface{}, logChannel chan string) bool {
	var path string = fmt.Sprintf("%v", arguments[0])
	skip, _ := strconv.Atoi(fmt.Sprintf("%v", arguments[1]))
	defer func() {
		// recover from panic caused by writing to a closed channel
		if r := recover(); r != nil {
		}
	}()
	action.InProgress = true
	logChannel <- fmt.Sprintf("Counting rows from file : %s", path)
	var file *os.File
	var err error
	if file, err = os.Open(path); err == nil {
		defer file.Close()
		logChannel <- fmt.Sprintf("File : %s found ...", path)
		var numRows int = 0
		scanner := bufio.NewScanner(file)
		if skip > 0 {
			logChannel <- fmt.Sprintf("Skipping : %d row(s) ...", skip)
			for i := 0; i < skip; i++ {
				scanner.Scan()
				scanner.Text()
			}
		}
		for scanner.Scan() {
			scanner.Text()
			numRows++
		}
		logChannel <- fmt.Sprintf("Number of Rows : %d, found in file %s", numRows, path)
		action.Message = fmt.Sprintf("File %s contains %d rows.", path, numRows)
		action.Success = true
		action.InProgress = false
		defer func() {
			// recover from panic caused by writing to a closed channel
			if r := recover(); r != nil {
			}
		}()
		return true
	}
	logChannel <- err.Error()
	action.Message = fmt.Sprintf("Error in count file rows : %s", err.Error())
	action.Success = false
	action.InProgress = false
	return false
}

type CounterParser struct {
	Skip int
	Path string
}

func (p *CounterParser) Validate() bool {
	flag.StringVar(&p.Path, "file", "", "Full qualified file path")
	flag.IntVar(&p.Skip, "skip", 0, "Number of row to skip count from file top")
	flag.Parse()
	return p.Path != ""
}

func (p *CounterParser) Parse() []interface{} {
	var arguments []interface{} = make([]interface{}, 0)
	arguments = append(arguments, p.Path)
	arguments = append(arguments, p.Skip)
	return arguments
}
