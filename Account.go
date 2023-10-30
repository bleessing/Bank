package Bank

import (
	"github.com/google/uuid"
)

type Account struct {
	ID      uuid.UUID
	Active  bool
	Balance float64
	Limit   float64
	Cards   map[uuid.UUID]*Card
}

func CreateAccount() *Account {
	account := &Account{
		ID:     uuid.Must(uuid.NewUUID()),
		Active: true,
		Cards:  make(map[uuid.UUID]*Card),
	}
	account.AddCard(CreateCard())
	return account
}

func (a *Account) AddCard(c *Card) {
	a.Cards[c.ID] = c
}
func (a *Account) CloseCard(c *Card) {
	a.Cards[c.ID] = nil
}
func (a *Account) ViewBalance() float64 {
	return a.Balance
}
