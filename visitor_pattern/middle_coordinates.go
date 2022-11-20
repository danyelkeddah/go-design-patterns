package visitor_pattern

import "fmt"

type MiddleCoordinates struct {
	x int
	y int
}

func (c *MiddleCoordinates) VisitForSquare(s *Square) {
	fmt.Println("calculating middle point for square")
}

func (c *MiddleCoordinates) VisitForCircle(s *Circle) {
	fmt.Println("calculating middle point for circle")
}

func (c *MiddleCoordinates) VisitForRectangle(s *Rectangle) {
	fmt.Println("calculating middle point for rectangle")
}
