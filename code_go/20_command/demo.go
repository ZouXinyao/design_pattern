package main

import "github.com/ZouXinyao/design_pattern/code_go/20_command/command"

func main() {
	executor := command.NewExecutor()

	runCommand := command.NewRunCommand(executor)
	stopCommand := command.NewStopCommand(executor)

	runDictator := command.NewDictator(runCommand)
	runDictator.Request()

	stopDictator := command.NewDictator(stopCommand)
	stopDictator.Request()
}