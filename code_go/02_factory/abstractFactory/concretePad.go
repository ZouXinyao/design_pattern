package abstractFactory

import "fmt"

type dellPad struct {
	brand string
}

func (pc *dellPad) PrintProductBrand() {
	fmt.Println("A pad of " + pc.brand)
}

type lenovoPad struct {
	brand string
}

func (pc *lenovoPad) PrintProductBrand() {
	fmt.Println("A pad of " + pc.brand)
}