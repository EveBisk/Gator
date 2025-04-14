package domain

import (
	"fmt"

	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID
	Name string
}

func (u User) Print() {
	fmt.Printf(" * ID:      %v\n", u.ID)
	fmt.Printf(" * Name:    %v\n", u.Name)
}
