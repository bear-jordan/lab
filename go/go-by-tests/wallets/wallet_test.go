package wallet

import "testing"

func assertBalance(t *testing.T, got Bitcoin, want Bitcoin) {
    t.Helper()

    if got != want {
        t.Errorf("got %s, want %s", got, want)
    }
}

func assertNoError(t *testing.T, err error) {
    t.Helper()

    if err != nil {
        t.Errorf("error found when there should not be one")
    }
}

func assertError(t *testing.T, got error, want error) {
    t.Helper()

    if got == nil {
        t.Fatalf("expected an error, got nil")
    }

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}

func TestWallet(t *testing.T) {
    t.Run("checking deposit", func(t *testing.T) {
        wallet := Wallet{}

        wallet.Deposit(Bitcoin(10))

        got := wallet.Balance()
        want := Bitcoin(10)

        assertBalance(t, got, want)

    })

    t.Run("checking withdrawl", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        err := wallet.Withdrawl(Bitcoin(5))

        got := wallet.Balance()
        want := Bitcoin(5)

        assertNoError(t, err)
        assertBalance(t, got, want)
    })

    t.Run("overdraft balance, should error", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))
        err := wallet.Withdrawl(Bitcoin(100))
        got := wallet.Balance()
        want := Bitcoin(10)

        assertError(t, err, ErrOverdraft)
        assertBalance(t, got, want)
    })
}
