package main

import (
	"fmt"
	"github.com/ZouXinyao/design_pattern/code_go/18_visitor/visitor"
)

func main() {
	s := visitor.NewSquare(10)
	c := visitor.NewCircle(10)
	r := visitor.NewRectangle(5, 10)

	aVisitor := visitor.NewAreaVisitor()
	pVisitor := visitor.NewPerimeterVisitor()

	s.Accept(aVisitor)
	fmt.Println("Area of square is ", aVisitor.GetArea())
	c.Accept(aVisitor)
	fmt.Println("Area of circle is ", aVisitor.GetArea())
	r.Accept(aVisitor)
	fmt.Println("Area of rectangle is ", aVisitor.GetArea())

	s.Accept(pVisitor)
	fmt.Println("Perimeter of square is ", pVisitor.GetPerimeter())
	c.Accept(pVisitor)
	fmt.Println("Perimeter of circle is ", pVisitor.GetPerimeter())
	r.Accept(pVisitor)
	fmt.Println("Perimeter of rectangle is ", pVisitor.GetPerimeter())


}
