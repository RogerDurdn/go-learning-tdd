package pointers_error

import (
	"errors"
	"fmt"
)

var ErrInsufficientBalance = errors.New("insufficient balance to withdraw")

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BCT", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientBalance
	}
	w.balance -= amount
	return nil
}
