package main

import "github.com/ZouXinyao/design_pattern/code_go/21_mediator/mediator"

func main() {
	reqMediator := mediator.NewMediator()

	req01 := mediator.NewLowPriorityReq(1, reqMediator)
	req02 := mediator.NewLowPriorityReq(2, reqMediator)
	req03 := mediator.NewHighPriorityReq(3, reqMediator)

	req01.Process()
	req02.Process()
	req03.Process()

	req01.Release()

}


