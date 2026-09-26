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

type Transaction struct {
	FromID string
	ToID   string
	Amount float64
}

type PaymentSystem struct {
	Users            map[string]*User
	TransactionQueue []Transaction
}

func (p *PaymentSystem) AddUser(user *User) {
	p.Users[user.ID] = user
}

func (p *PaymentSystem) AddTransaction(t Transaction) {
	p.TransactionQueue = append(p.TransactionQueue, t)
}

func (p *PaymentSystem) ProcessingTransactions(t Transaction) error {
	user1, ok1 := p.Users[t.FromID]
	user2, ok2 := p.Users[t.ToID]

	if !ok1 || !ok2 {
		return errors.New("User not found")
	}

	if err := user1.Withdraw(t.Amount); err != nil {
		return err
	}

	user2.Deposit(t.Amount)

	return nil
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
	t1 := Transaction{FromID: "1", ToID: "2", Amount: 200}
	t2 := Transaction{FromID: "2", ToID: "1", Amount: 50}
	paysys := &PaymentSystem{
		Users:            make(map[string]*User),
		TransactionQueue: make([]Transaction, 0),
	}
	paysys.AddUser(user1)
	paysys.AddUser(user2)
	paysys.AddTransaction(t1)
	paysys.AddTransaction(t2)

	for id, i := range paysys.TransactionQueue {
		if err := paysys.ProcessingTransactions(i); err != nil {
			fmt.Println("Error: ", err)
			continue
		} else {
			fmt.Printf("Success Transaction %d, Balance y1:  %.0f, Balance y2:  %.0f\n", id+1, user1.Balance, user2.Balance)
		}
	}
}
