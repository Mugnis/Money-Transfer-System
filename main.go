package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID      string
	Name    string
	Balance float64
}

func (u *User) Deposit(amount float64) {
	u.Balance += amount
}

func (u *User) Withdraw(amount float64) error {
	if u.Balance < amount {
		return errors.New("Insufficient funds")
	} else {
		u.Balance -= amount
		return nil
	}
}

func main() {
	user1 := &User{ID: "1", Name: "User1", Balance: 1000}
	user2 := &User{ID: "2", Name: "User2", Balance: 500}
	user3 := &User{ID: "3", Name: "User3", Balance: 10}
	user1.Withdraw(10)
	fmt.Println(user1.Balance)
	user2.Deposit(100)
	fmt.Println(user2.Balance)
	user3.Withdraw(100)
	fmt.Println(user3.Balance)
	user1.Deposit(40)
	fmt.Println(user1.Balance)
}
