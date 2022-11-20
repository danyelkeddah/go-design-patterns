package builder

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (nb *NormalBuilder) SetWindowType() {
	nb.windowType = "Wooden window"
}

func (nb *NormalBuilder) SetDoorType() {
	nb.doorType = "Wooden door"
}

func (nb *NormalBuilder) SetNumFloor() {
	nb.floor = 2
}

func (nb *NormalBuilder) GetHouse() House {
	return House{
		doorType:   nb.doorType,
		windowType: nb.windowType,
		floor:      nb.floor,
	}
}
