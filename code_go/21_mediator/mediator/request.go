package mediator

import (
	"fmt"
)

type Request interface {
	Process()
	Release()
	isHighPriority() bool
}

type ConcreteReq struct {
	reqId    int
	priority bool
	mediator Mediator
}

func NewHighPriorityReq(id int, mediator Mediator) *ConcreteReq{
	return &ConcreteReq{
		reqId: id,
		priority: true,
		mediator: mediator,
	}
}

func NewLowPriorityReq(id int, mediator Mediator) *ConcreteReq{
	return &ConcreteReq{
		reqId: id,
		priority: false,
		mediator: mediator,
	}
}

func (req *ConcreteReq) Process() {
	if !req.mediator.processingReq(req) {
		fmt.Println("sys is busy, pls wait!", req.reqId)
		return
	}
	fmt.Println("Processing the current req", req.reqId)
}

func (req *ConcreteReq) Release() {
	fmt.Println("release the current req", req.reqId)
	req.mediator.notifyAboutRelease()
}

func (req *ConcreteReq) isHighPriority() bool {
	return req.priority
}



