package bridge

type commonSpeed struct {
	*AbstractMessage
}

func (msg *commonSpeed) sendMsg() {
	msg.AbstractMessage.sendMsg()
}

func NewCommonMsg(s Sending) *commonSpeed {
	return &commonSpeed{
		&AbstractMessage{
			spendTime: 5,
			sendCompany: s,
		},
	}
}
