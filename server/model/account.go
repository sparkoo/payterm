package model

import (
	"fmt"
	"math/rand"
)

type UserId string

type Account struct {
	id      UserId
	name    string
	balance int
}

func (a *Account) Withdraw(amount int) int {
	a.balance = a.balance - amount
	return a.balance
}

func (a *Account) Balance() int {
	return a.balance
}

func (a *Account) Name() string {
	return a.name
}

func (a *Account) Id() UserId {
	return a.id
}

func NewAccount(name string, balance int) *Account {
	return &Account{
		id:      UserId(fmt.Sprintf("%v", rand.Uint64())),
		name:    name,
		balance: balance,
	}
}