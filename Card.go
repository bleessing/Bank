package Bank

import (
	"fmt"
	"github.com/google/uuid"
)

type Card struct {
	ID     uuid.UUID
	Active bool
	Block  bool
}

func CreateCard() *Card {
	return &Card{
		ID: uuid.Must(uuid.NewUUID()),
	}
}
func (c *Card) ActivateCard() {
	if c.Active != true {
		c.Active = true
	}
}
func (c *Card) BlockedCard() {
	if c.Active != false {
		c.Block = true
	}
}
func (c *Card) ViewStatusCard() string {
	return fmt.Sprintf("Status is: Active: %v, Blocked: %v", c.Active, c.Block)
}
