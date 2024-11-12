package account

import "sync"

// Define the Account type here.
type Account struct {
	amount int64
	isOpen bool
	mu     sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	account := Account{amount: amount, isOpen: true, mu: sync.Mutex{}}
	return &account
}

func (a *Account) Balance() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.isOpen {
		return 0, false
	}
	return a.amount, a.isOpen
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.isOpen {
		return 0, false
	}
	if a.amount+amount < 0 {
		return a.amount, false
	}
	a.amount += amount
	return a.amount, true
}

func (a *Account) Close() (int64, bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if !a.isOpen {
		return 0, false
	}
	ammount := a.amount
	a.amount = 0
	a.isOpen = false
	return ammount, true
}
