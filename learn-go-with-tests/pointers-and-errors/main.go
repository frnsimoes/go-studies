package main

import "fmt"

type Bitcoin int

type Stringer interface {
}
type Wallet struct {
	balance Bitcoin
}

// pointer to Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	// return (*w).balance
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
