package pointers_error_test

import (
	w "projectsBook/goTDD/pointers-errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := w.Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, w.Bitcoin(10))
	})

	t.Run("Withdraw with founds", func(t *testing.T) {
		wallet := w.Wallet{}
		wallet.Deposit(40)
		err := wallet.Withdraw(w.Bitcoin(20))

		assertNoError(t, err)
		assertBalance(t, wallet, w.Bitcoin(20))
	})

	t.Run("Withdraw insufficient balance", func(t *testing.T) {
		startingBalance := w.Bitcoin(20)
		wallet := w.Wallet{}
		wallet.Deposit(startingBalance)
		err := wallet.Withdraw(w.Bitcoin(100))

		assertError(t, err, w.ErrInsufficientBalance)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet w.Wallet, want w.Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got != want {
		t.Errorf("got %s, want %s", got.Error(), want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("Wanted no error, but got err: ", err)
	}
}
