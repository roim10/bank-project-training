package main

import "errors"

var (
	ErrInsufficientFunds    = errors.New("insufficient funds")
	ErrAccountNotFound      = errors.New("account not found")
	ErrInvalidAmount        = errors.New("invalid amount")
	ErrAccountAlreadyExists = errors.New("account already exists")
)
