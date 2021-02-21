package facade

import "fmt"

type shape interface {
	draw()
}

type rectangle struct {

}

func (rec *rectangle) draw() {
	fmt.Println("rectangle draw()")
}

type circle struct {

}

func(cir *circle) draw() {
	fmt.Println("circle draw()")
}

type ShapeMaker struct {
	*circle
	*rectangle
}

func NewShapeMaker() *ShapeMaker {
	return &ShapeMaker{
		&circle{},
		&rectangle{},
		}
}

func (maker *ShapeMaker) DrawCircle() {
	maker.circle.draw()
}

func (maker *ShapeMaker) DrawRectangle() {
	maker.rectangle.draw()
}