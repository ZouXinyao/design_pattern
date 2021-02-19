package bridge

type AbstractMessage struct {
	spendTime   int
	sendCompany Sending
}

func (msg *AbstractMessage) sendMsg() {
	msg.sendCompany.Send()
}