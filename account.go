package main

import "sync"

type Transactor interface {
	Deposit(amount float64) (float64, error)
	Withdraw(amount float64) (float64, error)
}

type Account struct {
	ID      string
	Owner   string
	balance float64
	mu      sync.Mutex
}

func (acc *Account) Deposit(amount float64) (float64, error) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if amount <= 0 {
		return 0, ErrInvalidAmount
	}
	acc.balance += amount
	return acc.balance, nil
}

func (acc *Account) Withdraw(amount float64) (float64, error) {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	if amount <= 0 {
		return 0, ErrInvalidAmount
	}
	if amount > acc.balance {
		return 0, ErrInsufficientFunds
	}
	acc.balance -= amount
	return acc.balance, nil
}
func (acc *Account) GetBalance() float64 {
	acc.mu.Lock()
	defer acc.mu.Unlock()
	return acc.balance
}
