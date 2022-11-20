package decorator

type VeggeMania struct{}

func (v *VeggeMania) GetPrice() int {
	return 15
}
