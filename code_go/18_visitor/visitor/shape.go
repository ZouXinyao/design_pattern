package visitor

type Shape interface {
	Accept(Visitor)
}

type Square struct {
	side int
}

func NewSquare(side int) *Square {
	return &Square{
		side: side,
	}
}

func (s *Square) Accept(v Visitor) {
	v.visitForSquare(s)
}

type Circle struct {
	radius int
}

func NewCircle(radius int) *Circle {
	return &Circle{
		radius: radius,
	}
}

func (c *Circle) Accept(v Visitor) {
	v.visitForCircle(c)
}

type Rectangle struct {
	length int
	width  int
}

func NewRectangle(length, width int) *Rectangle {
	return &Rectangle{
		length: length,
		width: width,
	}
}

func (r *Rectangle) Accept(v Visitor) {
	v.visitForRectangle(r)
}


