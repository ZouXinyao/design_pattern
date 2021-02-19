package bridge

import "fmt"

type Yunda struct {
	infor   *baseInfor
	content *contentMsg
}

func (y *Yunda) Send() {
	fmt.Println(y.content.msg + " to " + y.infor.receiver)
}

func (y *Yunda) SetSenderAndReceiver(sender string, receiver string) {
	y.infor.sender = sender
	y.infor.receiver = receiver
}

func NewYundaSender(msg string) *Yunda {
	y := Yunda{}
	y.content.msg = msg
	return &y
}