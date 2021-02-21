package template

import "fmt"

type Message interface {
	ReadMassage() string
	SaveMassage(msg string)
	SendMassage(receiver string)
}

type QQ struct {
	msg string
}

func (q *QQ) SaveMassage(msg string) {
	q.msg = msg
}

func (q *QQ) ReadMassage() string {
	return q.msg
}

func (q *QQ) SendMassage(receiver string) {
	fmt.Println("Sending QQ msg to " + receiver)
	fmt.Println("The message is " + q.msg)
}

func NewQQMessage() Message{
	return &QQ{}
}

type Wechat struct {
	msg string
}

func (w *Wechat) SaveMassage(msg string) {
	w.msg = msg
}

func (w *Wechat) ReadMassage() string {
	return w.msg
}

func (w *Wechat) SendMassage(receiver string) {
	fmt.Println("Sending wechat msg to " + receiver)
	fmt.Println("The message is " + w.msg)
}

func NewWechatMessage() Message{
	return &Wechat{}
}

