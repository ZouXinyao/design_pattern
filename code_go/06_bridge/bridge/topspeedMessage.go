package bridge

type topSpeed struct {
	*AbstractMessage
}

func (msg *topSpeed) sendMsg() {
	msg.AbstractMessage.sendMsg()
}

func NewTopSpeedMsg(s Sending) *topSpeed {
	return &topSpeed{
		&AbstractMessage{
			spendTime: 2,
			sendCompany: s,
		},
	}
}