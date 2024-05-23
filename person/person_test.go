package person

import (
	"testing"
)

func TestPersonString(t *testing.T) {
	tests := []struct {
		person   Person
		expected string
	}{
		{
			Person{Name: "Alice", InitialContribution: 50, HasToPay: -20, MaxAfford: 100},
			"Alice -> -20.00 (Will collect)",
		},
		{
			Person{Name: "Bob", InitialContribution: 30, HasToPay: 20, MaxAfford: 50},
			"Bob -> 20.00 (Max Afford)",
		},
		{
			Person{Name: "Charlie", InitialContribution: 10, HasToPay: 40, MaxAfford: 100},
			"Charlie -> 40.00",
		},
		{
			Person{Name: "Diana", InitialContribution: 20, HasToPay: 0, MaxAfford: 20},
			"Diana -> 0.00 (Max Afford)",
		},
		{
			Person{Name: "Eve", InitialContribution: 0, HasToPay: 50, MaxAfford: 100},
			"Eve -> 50.00",
		},
	}

	for _, test := range tests {
		result := test.person.String()
		if result != test.expected {
			t.Errorf("Expected %s but got %s", test.expected, result)
		}
	}
}
