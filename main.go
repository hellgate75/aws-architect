package main

import (
	"aws-architect/bootstrap"
	"aws-architect/abstract"
	"fmt"
	"os"
	"strings"
	"time"
	"aws-architect/log"
)

var currentAction abstract.ActionOps
const loggingTimeout = 10
var logger log.Logger

func PrintHelp(command string) () {
	var size int = abstract.GetActionRegistry().Size()
	for i := 0; i < size; i++ {
		var ops abstract.ActionOps = abstract.ActiveActionRegistry.ElementAt(i)
		if ops.GetCommand() == command {
			println("aws-architects", ops.GetUsage())
			return
		}
	}
	PrintUsage("")
}


func PrintUsage(command string) () {
	if strings.ToLower(command) == "help" && len(os.Args) > 2 {
		PrintHelp(strings.ToLower(os.Args[2]))
	} else {
		var size int = abstract.GetActionRegistry().Size()
		println("aws-architect <command> parameters ...")
		println("Type command : help <commnand> to receive more information")
		println("Available commands : ")
		for i := 0; i < size; i++ {
			var ops abstract.ActionOps = abstract.ActiveActionRegistry.ElementAt(i)
			println("\t", ops.GetCommand())
		}
	}
}

func init() {
	bootstrap.InitActions()
	logger = log.Logger{
		NoDebug: bootstrap.Settings.DebugDisabled,
	}
	var size int = abstract.GetActionRegistry().Size()
	logger.Debug(fmt.Sprintf("Available Action %d", size))
	var found bool = false
	if len(os.Args) > 1 {
		for i := 0; i < size; i++ {
			var ops abstract.ActionOps = abstract.ActiveActionRegistry.ElementAt(i)
			if strings.ToLower(ops.GetCommand()) == strings.ToLower(os.Args[1]) {
				found = true
				currentAction = ops
				break
			}
		}
	}
	if ! found {
		if len(os.Args) > 1 {
			PrintUsage(os.Args[1])
		} else {
			PrintUsage("")
		}
		os.Exit(1)
	}
}

func main() {
	var startTime time.Time = time.Now()
	var args []string = make([]string, 0)
	args = append(args, os.Args[0])
	args = append(args, os.Args[2:]...)
	os.Args=args
	var satisfied bool = currentAction.AcquireValues()
	logger.Log(fmt.Sprintf("current action : %s", strings.ToLower(currentAction.GetCommand())))
	if satisfied {
		var logChannel chan string = make(chan string)
		var response bool = false
		go func(channel chan  string, progress func ()(bool)) () {
			time.Sleep(time.Millisecond*500)
			for progress() {
				select {
				case message := <-channel:
					if progress() {
						go logger.Log(message)
					}
				case <-time.After(time.Second * loggingTimeout):
					logger.Debug("Logger timeout ...")
				}
			}
		}(logChannel, currentAction.IsInProgress)
		response=currentAction.Execute(logChannel)
		logger.Log(fmt.Sprintf("Execution of command %s complated : %t", currentAction.GetCommand(), response))
		logger.Log(fmt.Sprintf("Exit Code : %d", currentAction.GetExitCode()))
		logger.Log(fmt.Sprintf("Message : %s", currentAction.GetLastMessage()))
		close(logChannel)
	} else  {
		var args []string = make([]string, 0)
		args = append(args, os.Args[0])
		args = append(args, "help")
		args = append(args, currentAction.GetCommand())
		os.Args=args
		println("Command :", currentAction.GetCommand(),"-> wrong parameters")
		PrintUsage("help")
	}
	duration := time.Since(startTime)
	print("Done!!\n")
	print(fmt.Sprintf("Total time : %v\n", duration))
}
