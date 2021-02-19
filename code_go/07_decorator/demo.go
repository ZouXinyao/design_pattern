package main

import "github.com/ZouXinyao/design_pattern/code_go/07_decorator/decorator"

func main() {
	circle := decorator.NewCircle()
	rectangle := decorator.NewRectangle()

	redCircle := decorator.NewRedShapeDecorator(circle)
	redRectangle := decorator.NewRedShapeDecorator(rectangle)

	circle.Draw()

	redCircle.Draw()

	redRectangle.Draw()

}