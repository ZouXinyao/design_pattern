package main

import "github.com/ZouXinyao/design_pattern/code_go/13_template/template"

func main() {
	Msg := template.NewQQMessage()
	Msg.SaveMassage("a qq msg")
	Msg.ReadMassage()
	Msg.SendMassage("Mary")

	Msg = template.NewWechatMessage()
	Msg.SaveMassage("a wechat msg")
	Msg.ReadMassage()
	Msg.SendMassage("Tom")
}
