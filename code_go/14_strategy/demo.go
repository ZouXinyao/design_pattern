package main

import (
	"fmt"
	"github.com/ZouXinyao/design_pattern/code_go/14_strategy/strategy"
)

func main() {
	context := strategy.NewContext(strategy.NewAdd())
	fmt.Println(context.ExecuteStrategy(5, 6))

	context = strategy.NewContext(strategy.NewSubtract())
	fmt.Println(context.ExecuteStrategy(5, 6))

	context = strategy.NewContext(strategy.NewMultiply())
	fmt.Println(context.ExecuteStrategy(5, 6))


}
