package examples

import "github.com/google/uuid"

type user struct {
	ID        string
	Email     string
	FirstName string
}

func DefaultUser() user {
	return user{
		ID:        uuid.NewString(),
		Email:     "example@pvarber.com",
		FirstName: "Pablo",
	}
}
