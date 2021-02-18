package main

import (
	"My_golang_design_pattern/03_builder/builder"
	"fmt"
)

func main() {

	dellBuilder := &builder.DellBuilder{}
	dellDirector := builder.NewDirector(dellBuilder)
	dellDirector.Construct("i7-9700k", "RTX2060", "z370", "EVGA650W")
	fmt.Println(dellBuilder.GetResult())

	lenovoBuilder := &builder.LenovoBuilder{}
	lenovoDirector := builder.NewDirector(lenovoBuilder)
	lenovoDirector.Construct("i7-10700k", "RTX3060", "z470", "EVGA800W")
	fmt.Println(lenovoBuilder.GetResult())
}
