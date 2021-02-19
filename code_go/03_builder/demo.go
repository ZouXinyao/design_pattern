package main

import (
	"fmt"
	"github.com/ZouXinyao/design_pattern/code_go/03_builder/builder"
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
