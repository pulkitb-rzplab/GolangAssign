package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account struct {
	mu      sync.Mutex
	balance float64
}

func (a *Account) Deposit(amount float64) {
	a.mu.Lock()
	time.Sleep(time.Second * 2)
	defer a.mu.Unlock()
	a.balance += amount
	fmt.Println("Deposited: ", amount, " at time : ", time.Now())
}
func (a *Account) withdraw(amount float64) {
	a.mu.Lock()
	time.Sleep(time.Second * 2)
	if a.balance < amount {
		fmt.Println("Insufficient balance")
		a.mu.Unlock()
		return
	}
	defer a.mu.Unlock()
	a.balance -= amount
	fmt.Println("Withdrawn: ", amount, " at time : ", time.Now())
}

func main() {
	acc := Account{balance: 500}
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var wg sync.WaitGroup

	update := func(typ int) {
		for range 25 {
			amt := (r.Float64() * 900) + 100
			if typ == 0 {
				acc.Deposit(amt)
			} else {
				acc.withdraw(amt)
			}

		}
		wg.Done()

	}

	wg.Add(2)
	go update(0)
	go update(1)

	wg.Wait()
	fmt.Println("Final balance: ", acc.balance)

}
