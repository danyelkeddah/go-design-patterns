package builder

type Builder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() House
}
