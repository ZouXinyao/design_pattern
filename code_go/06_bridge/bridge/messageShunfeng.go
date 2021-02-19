package bridge

import "fmt"

type Shunfeng struct {
	infor   *baseInfor
	content *contentMsg
}

func (s *Shunfeng) Send() {
	fmt.Println(s.content.msg + " to " + s.infor.receiver)
}

func (s *Shunfeng) SetSenderAndReceiver(sender string, receiver string) {
	s.infor.sender = sender
	s.infor.receiver = receiver
}

func NewShunfengSender(msg string) *Shunfeng {
	s := Shunfeng{}
	s.content.msg = msg
	return &s
}
