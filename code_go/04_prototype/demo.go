package main

import (
	"fmt"
	"github.com/ZouXinyao/design_pattern/code_go/04_prototype/prototype"
)

func main() {
	intelMainFrame := prototype.NewMainFrame()
	intelMainFrame.SetCpu("i7-10700k")
	intelMainFrame.SetGpu("RTX2070s")
	dellComputer := prototype.NewComputer("Dell")
	dellComputer.SetScreen("ASUS")
	dellComputer.SetMainframe(intelMainFrame)

	lenovoComputer := dellComputer.Clone()
	lenovoComputer.SetBrand("Lenovo")

	fmt.Println("Dell:")
	fmt.Println(dellComputer.GetBrand())
	fmt.Println(dellComputer.GetScreen())
	fmt.Println(dellComputer.Pc.GetCpu())
	fmt.Println(dellComputer.Pc.GetGpu())

	fmt.Println("Lenovo:")
	fmt.Println(lenovoComputer.GetBrand())
	fmt.Println(lenovoComputer.GetScreen())
	fmt.Println(lenovoComputer.Pc.GetCpu())
	fmt.Println(lenovoComputer.Pc.GetGpu())


}
