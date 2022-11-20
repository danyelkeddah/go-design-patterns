package observer

type Observer interface {
	Update(string)
	GetId() string
}
