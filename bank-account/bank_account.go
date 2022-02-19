package account

import "sync"

type Account struct {sync.Mutex; open bool; balance int64}

func Open(amount int64) *Account {
    if amount < 0 { return nil }
    return &Account{open: true, balance: amount}
}

func (a *Account) Balance() (int64, bool) {
    if a.open == false { return 0, false }
    return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
    a.Lock()
    defer a.Unlock()
    if a.open == false || a.balance < -amount { return 0, false }
    a.balance += amount
    return a.balance, true
}

func (a *Account) Close() (int64, bool) {
    a.Lock()
	defer a.Unlock()
	if a.open == false { return 0, false }
	a.open = false
	return a.balance, true
}

