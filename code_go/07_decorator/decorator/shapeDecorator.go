package decorator

import "fmt"

type ShapeDecorator struct {
	Shape
}

func (sd *ShapeDecorator) Draw() {
	sd.Shape.Draw()
}

func NewShapeDecorator(s Shape) *ShapeDecorator {
	return &ShapeDecorator{Shape: s}
}

type RedShapeDecorator struct {
	ShapeDecorator
}

func (rsd *RedShapeDecorator) Draw() {
	rsd.Shape.Draw()
	rsd.setRedBorder(rsd.Shape)
}

func (rsd *RedShapeDecorator) setRedBorder(s Shape) {
	fmt.Println("Border Color: Red")
}

func NewRedShapeDecorator(s Shape) *RedShapeDecorator {
	return &RedShapeDecorator{ShapeDecorator{Shape: s}}
}