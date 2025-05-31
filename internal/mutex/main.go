package main

import "sync"

type Account struct {
	mu      sync.RWMutex
	balance int
}

func (account *Account) Deposit(amount int) {
	account.mu.Lock()
	defer account.mu.Unlock()
	account.balance += amount

}

func (account *Account) GetBalance() int {
	account.mu.RLock()
	defer account.mu.RUnlock()

	return account.balance
}
