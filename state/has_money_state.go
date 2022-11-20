package state

import "fmt"

type HasMoneyState struct {
	VendingMachine *VendingMachine
}

func (s HasMoneyState) AddItem(i int) error {
	return fmt.Errorf("item dispense in progress")
}

func (s HasMoneyState) RequestItem() error {
	return fmt.Errorf("item dispense in progress")
}

func (s HasMoneyState) InsertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

func (s HasMoneyState) DispenseItem() error {
	fmt.Println("dispensing item")
	s.VendingMachine.ItemCount = s.VendingMachine.ItemCount - 1
	if s.VendingMachine.ItemCount == 0 {
		s.VendingMachine.SetState(s.VendingMachine.NoItem)
	} else {
		s.VendingMachine.SetState(s.VendingMachine.HasItem)
	}
	return nil
}
