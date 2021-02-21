package memento

type Message struct {
	content string
}

func NewMessage() *Message {
	return &Message{}
}

func (msg *Message) CreatMemento() *Memento {
	return &Memento{
		content: msg.content,
	}
}

func (msg *Message) Restore(mem *Memento) {
	msg.content = mem.GetSavedContent()
}

func (msg *Message) FillMsg(content string) {
	msg.content = content
}

func (msg *Message) ReadMsg() string {
	return msg.content
}

