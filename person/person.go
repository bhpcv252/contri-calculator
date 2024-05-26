package person

import (
	"fmt"
)

type Person struct {
	Name                string
	InitialContribution float32
	HasToPay            float32
	AffordLimit         float32
	CanAfford           float32
}

func (p Person) String() string {
	if p.HasToPay < 0 {
		// This person will collect money because they have paid more than their share.
		return fmt.Sprintf("%s -> %.2f (Will collect)", p.Name, p.HasToPay)
	}

	if p.CanAfford == 0 {
		// This person can only pay the maximum amount they can afford.
		return fmt.Sprintf("%s -> %.2f (Max Afford)", p.Name, p.HasToPay)
	}

	return fmt.Sprintf("%s -> %.2f", p.Name, p.HasToPay)
}
