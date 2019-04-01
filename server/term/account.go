package term

import (
	"fmt"
	"math/rand"
)

type Account struct {
	id      userid
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

func NewAccount(name string, balance int) *Account {
	return &Account{
		id:      userid(fmt.Sprintf("%v", rand.Uint64())),
		name:    name,
		balance: balance,
	}
}
