package command

type Command interface {
	execute()
}

type RunCommand struct {
	executor Executor
}

func NewRunCommand(e Executor) *RunCommand {
	return &RunCommand{
		executor: e,
	}
}

func (r *RunCommand) execute() {
	r.executor.run()
}

type StopCommand struct {
	executor Executor
}

func NewStopCommand(e Executor) *StopCommand {
	return &StopCommand{
		executor: e,
	}
}

func (r *StopCommand) execute() {
	r.executor.stop()
}