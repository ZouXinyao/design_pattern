package visitor

import "math"

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

type AreaVisitor struct {
	area float64
}

func NewAreaVisitor() *AreaVisitor {
	return &AreaVisitor{}
}

func (a *AreaVisitor) GetArea() float64 {
	return a.area
}

func (a *AreaVisitor) visitForSquare(s *Square) {
	a.area = float64(s.side * s.side)
}

func (a *AreaVisitor) visitForCircle(c *Circle) {
	a.area = math.Pi * float64(c.radius * c.radius)
}

func (a *AreaVisitor) visitForRectangle(r *Rectangle) {
	a.area = float64(r.length * r.width)
}

type PerimeterVisitor struct {
	perimeter float64
}

func NewPerimeterVisitor() *PerimeterVisitor {
	return &PerimeterVisitor{}
}

func (p *PerimeterVisitor) GetPerimeter() float64 {
	return p.perimeter
}

func (p *PerimeterVisitor) visitForSquare(s *Square) {
	p.perimeter = float64(4 * s.side)
}

func (p *PerimeterVisitor) visitForCircle(c *Circle) {
	p.perimeter = 2 * math.Pi * float64(c.radius)
}

func (p *PerimeterVisitor) visitForRectangle(r *Rectangle) {
	p.perimeter = float64(2 * (r.length + r.width))
}