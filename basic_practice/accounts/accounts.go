package accounts

import (
	"errors"
	"fmt"
)

var errNoMoney = errors.New("Can't Withdraw")

// Account struct
type Account struct {
	// lowercase variables are private
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner}
	return &account
}

// reciever : first character should be first character of struct name
// if function has reciever --> method
// reciever is not original pointer, it is new copy without * (pointer operator)
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// balance of your account
// *Account : reciever pointer
func (a *Account) Balance() int {
	return a.balance
}

// Withdraw of your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Owner of the account
func (a *Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \n Has: ", a.Balance())
}
