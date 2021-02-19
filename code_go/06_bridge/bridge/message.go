package bridge

type Sending interface {
	Send()
	SetSenderAndReceiver(sender string, receiver string)
}

type baseInfor struct {
	sender   string
	receiver string
}

type contentMsg struct {
	msg string
}