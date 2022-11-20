package visitor_pattern

import "fmt"

type AreaCalculator struct {
	Area int
}

func (c *AreaCalculator) VisitForSquare(s *Square) {
	fmt.Println("calculating area for square")
}

func (c *AreaCalculator) VisitForCircle(s *Circle) {
	fmt.Println("calculating area for circle")
}

func (c *AreaCalculator) VisitForRectangle(s *Rectangle) {
	fmt.Println("calculating area for rectangle")
}
