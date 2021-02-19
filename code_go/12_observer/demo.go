package main

import (
	"github.com/ZouXinyao/design_pattern/code_go/12_observer/observer"
)

func main() {
	subject := observer.NewSubject()

	// 被观察者中注册观察者。
	reader1 := observer.NewReader("reader1")
	reader2 := observer.NewReader("reader2")
	reader3 := observer.NewReader("reader3")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.UpdateContext("observer mode")
	// Output:
	// reader1 receive observer mode
	// reader2 receive observer mode
	// reader3 receive observer mode
}
