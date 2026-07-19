package main

import (
	"fmt"
	"sync"
)

type BankStorage struct {
	Rmu     sync.RWMutex
	storage map[string]*Account
}

func (bs *BankStorage) GetAccount(key string) (*Account, error) {
	bs.Rmu.RLock()
	defer bs.Rmu.RUnlock()
	account, ok := bs.storage[key]
	if !ok {
		return nil, ErrAccountNotFound
	}
	return account, nil
}

func NewBankStorage() *BankStorage {
	return &BankStorage{
		storage: make(map[string]*Account),
	}
}

func (bs *BankStorage) AddAccount(account *Account) error {
	bs.Rmu.Lock()
	defer bs.Rmu.Unlock()
	_, exists := bs.storage[account.ID]
	if exists {
		return ErrAccountAlreadyExists
	}
	bs.storage[account.ID] = account
	return nil
}
func (bs *BankStorage) Transfer(FromID string, ToID string, amount float64) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("восстановились после паники: %v", r)
		}
	}()
	if amount < 0 {
		panic("the amount cannot be negative")
	}
	account1, err := bs.GetAccount(FromID)
	if err != nil {
		return err
	}
	account2, err := bs.GetAccount(ToID)
	if err != nil {
		return err
	}
	_, err = account1.Withdraw(amount)
	if err != nil {
		return err
	}
	_, err = account2.Deposit(amount)
	if err != nil {
		return err
	}
	return nil
}
