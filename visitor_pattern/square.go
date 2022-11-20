package visitor_pattern

type Square struct {
	side int
}

func (s *Square) Accept(v Visitor) {
	v.VisitForSquare(s)
}

func (s *Square) GetType() string {
	return "square"
}
