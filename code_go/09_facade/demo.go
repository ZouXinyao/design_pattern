package main

import "github.com/ZouXinyao/design_pattern/code_go/09_facade/facade"

func main() {
	maker := facade.NewShapeMaker()

	maker.DrawCircle()
	maker.DrawRectangle()
}
