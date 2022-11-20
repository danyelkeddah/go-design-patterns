package command

type OnCommand struct {
	Device Device
}

func (o *OnCommand) Execute() {
	o.Device.On()
}
