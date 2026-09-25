package main

import (
	"errors"
	"fmt"
	"sync"
)

type User struct {
	ID      string
	Name    string
	Balance float64
	mu      sync.Mutex
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
	user1, ok := p.Users[t.FromID]

	if !ok {
		return errors.New("User not found")
	}

	if err := user1.Withdraw(t.Amount); err != nil {
		return err
	}

	user2, ok := p.Users[t.ToID]

	if !ok {
		return errors.New("User not found")
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

func Worker(ch <-chan Transaction, ps *PaymentSystem, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range ch {
		if err := ps.ProcessingTransactions(t); err != nil {
			fmt.Println("Error: ", err)
			continue
		}
	}
}

func main() {
	user1 := &User{ID: "1", Name: "User1", Balance: 1000}
	user2 := &User{ID: "2", Name: "User2", Balance: 500}
	t1 := Transaction{FromID: "1", ToID: "2", Amount: 200}
	t2 := Transaction{FromID: "2", ToID: "1", Amount: 50}
	ps := &PaymentSystem{
		Users:            make(map[string]*User),
		TransactionQueue: make([]Transaction, 0),
	}
	ps.AddUser(user1)
	ps.AddUser(user2)
	ps.AddTransaction(t1)
	ps.AddTransaction(t2)
	ch := make(chan Transaction, len(ps.TransactionQueue))
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go Worker(ch, ps, &wg)
	}
	for _, t := range ps.TransactionQueue {
		ch <- t
	}
	close(ch)
	wg.Wait()
	fmt.Println("Итого")
	fmt.Printf("У первого пользователя должно получиться 850, а получилось %f\n", ps.Users["1"].Balance)
	fmt.Printf("У второго пользователя должно получиться 650, а получилось %f", ps.Users["2"].Balance)
}
