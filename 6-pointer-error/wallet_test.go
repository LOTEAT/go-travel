package main

import "testing"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(money int) {
	w.balance += money
}

func (w *Wallet) Balance() int {
	return (*w).balance
}

func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
