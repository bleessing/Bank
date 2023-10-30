package Bank

import "github.com/google/uuid"

type User struct {
	Id        uuid.UUID
	Firstname string
	Lastname  string
	Account   map[uuid.UUID]*Account
}

func CreateUSer() *User {
	return &User{
		Id:      uuid.Must(uuid.NewUUID()),
		Account: make(map[uuid.UUID]*Account),
	}
}
