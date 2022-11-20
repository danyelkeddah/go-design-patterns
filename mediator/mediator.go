package mediator

type Mediator interface {
	CanLand(Train) bool
	NotifyFree()
}
