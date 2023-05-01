package main

import "math/rand"

type Account struct {
	ID        int
	FirstName string
	LastName  string
	Number    int64
	Balance   int64
}

func NewAccount(firstname string, lastname string) Account {
	return Account{
		ID:        rand.Intn(10000),
		FirstName: firstname,
		LastName:  lastname,
		Number:    int64(rand.Intn(1000000)),
	}
}
