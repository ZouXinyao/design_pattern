package bridge

type urgentSpeed struct {
	*AbstractMessage
}

func (msg *urgentSpeed) sendMsg() {
	msg.AbstractMessage.sendMsg()
}

func NewUrgentMsg(s Sending) *urgentSpeed {
	return &urgentSpeed{
		&AbstractMessage{
			spendTime: 4,
			sendCompany: s,
		},
	}
}