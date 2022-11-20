package visitor_pattern

type Shape interface {
	GetType() string
	Accept()
}
