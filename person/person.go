package person

import (
	"fmt"
)

type Person struct {
	Name   string
	Amount float32
}

func (p Person) String() string {
	if p.Amount < 0 {
		return fmt.Sprintf("%s -> %.2f (Will collect)", p.Name, p.Amount)
	}
	return fmt.Sprintf("%s -> %.2f", p.Name, p.Amount)
}
