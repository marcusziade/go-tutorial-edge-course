package main

import "sync"

// Mutual exclusions

type Account struct {
	Balance int
	Mutex   sync.Mutex
}

func (a *Account) Withdraw(value int, waitGroup *sync.WaitGroup) {
	a.Mutex.Lock()
	a.Balance -= value
	a.Mutex.Unlock()
	waitGroup.Done()
}

func (a *Account) Deposit(value int, waitGroup *sync.WaitGroup) {
	a.Mutex.Lock()
	a.Balance += value
	a.Mutex.Unlock()
	waitGroup.Done()
}

func main() {
	account := Account{
		Balance: 1000,
	}

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go account.Withdraw(200, &waitGroup)
	go account.Deposit(5000, &waitGroup)
	waitGroup.Wait()

	println("Account Balance: ", account.Balance)
}
