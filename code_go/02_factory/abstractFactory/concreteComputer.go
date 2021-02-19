package abstractFactory

import "fmt"

type dellComputer struct {
	brand string
}

func (pc *dellComputer) PrintProductBrand() {
	fmt.Println("A computer of " + pc.brand)
}

type lenovoComputer struct {
	brand string
}

func (pc *lenovoComputer) PrintProductBrand() {
	fmt.Println("A computer of " + pc.brand)
}