package facade

import "fmt"

type Wallet struct {
	balance int
}

func newWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) CreditBalance(amount int) {
	w.balance += amount
	fmt.Println("wallet balance added successfully")
}

func (w *Wallet) DebitBalance(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("balance is not sufficient")
	}
	fmt.Printf("Wallet balance is sufficient")
	w.balance = w.balance - amount
	return nil
}
