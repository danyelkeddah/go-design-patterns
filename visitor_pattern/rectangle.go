package visitor_pattern

type Rectangle struct {
	l int
	b int
}

func (r *Rectangle) Accept(v Visitor) {
	v.VisitForRectangle(r)
}
func (r *Rectangle) GetType() string {
	return "rectangle"
}
