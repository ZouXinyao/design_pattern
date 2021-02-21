package command

import "fmt"

type Executor interface {
	run()
	stop()
}

type ConcreteExecutor struct {
	isRunning bool
}

func NewExecutor() *ConcreteExecutor {
	return &ConcreteExecutor{}
}

func (t *ConcreteExecutor) run() {
	t.isRunning = true
	fmt.Println("Executor run")
}

func (t *ConcreteExecutor) stop() {
	t.isRunning = false
	fmt.Println("Executor stop")
}