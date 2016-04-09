package account

import "sync"

type Account struct {
	m      sync.RWMutex
	amount int64
	closed bool
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{amount: initialDeposit}
}

func (a *Account) Close() (int64, bool) {
	a.m.Lock()
	defer a.m.Unlock()

	if a.closed {
		return 0, false
	}

	a.closed = true

	return a.amount, true
}

func (a *Account) Balance() (int64, bool) {
	a.m.RLock()
	defer a.m.RUnlock()

	if a.closed {
		return 0, false
	}

	return a.amount, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.m.Lock()
	defer a.m.Unlock()

	if a.closed {
		return 0, false
	}

	if a.amount+amount < 0 {
		return a.amount, false
	}

	a.amount += amount

	return a.amount, true
}