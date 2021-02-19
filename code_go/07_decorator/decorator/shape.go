package decorator

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct {
}

func (c *Circle) Draw() {
	fmt.Println("Shape: Circle")
}

func NewCircle() *Circle {
	return &Circle{}
}

type Rectangle struct {
}

func (r *Rectangle) Draw() {
	fmt.Println("Shape: Rectangle")
}

func NewRectangle() *Rectangle {
	return &Rectangle{}
}