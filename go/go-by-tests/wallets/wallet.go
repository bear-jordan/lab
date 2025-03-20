package wallet

import (
    "fmt"
    "errors"
)

var ErrOverdraft = errors.New("overdraft")

type Stringr interface {
    String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
    balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}

func (w *Wallet) Withdrawl(amount Bitcoin) error {
    if w.balance < amount {
        return ErrOverdraft
    }
    w.balance -= amount
    return nil
}

func (w Wallet) Balance() Bitcoin {
    return w.balance
}

