package Bank

import "github.com/google/uuid"

type Bank struct {
	Users map[uuid.UUID]*User
}

func (b *Bank) AddUser(u *User) {
	b.Users[u.Id] = u
}
func CreateBank() Bank {
	bank := Bank{
		Users: make(map[uuid.UUID]*User),
	}
	bank.AddUser(CreateUSer())
	return bank
}
