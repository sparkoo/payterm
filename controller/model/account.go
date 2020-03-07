package model

import "fmt"

type UserId string

type Account struct {
	id      UserId
	name    string
	balance int
}

func (a *Account) Withdraw(amount int) (int, error) {
	newBalance := a.balance - amount
	if newBalance < 0 {
		return 0, &NotEnoughMoneyOnAccountError{}
	}
	a.balance = newBalance
	return a.balance, nil
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

func NewAccount(id UserId, name string, balance int) *Account {
	return &Account{
		id:      id,
		name:    name,
		balance: balance,
	}
}

type NotEnoughMoneyOnAccountError struct {
}

func (err *NotEnoughMoneyOnAccountError) Error() string {
	return fmt.Sprintf("Not Enough money")
}
