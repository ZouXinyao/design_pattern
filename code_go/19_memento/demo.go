package main

import (
	"fmt"
	"github.com/ZouXinyao/design_pattern/code_go/19_memento/memento"
)

func main() {
	admin := memento.NewAdmin()
	msg01 := memento.NewMessage()

	msg01.FillMsg("0000001")
	admin.Add(msg01.CreatMemento())
	fmt.Println(msg01.ReadMsg())

	msg01.FillMsg("0000002")
	admin.Add(msg01.CreatMemento())
	fmt.Println(msg01.ReadMsg())

	msg01.Restore(admin.GetMemento(0))
	fmt.Println(msg01.ReadMsg())
}
