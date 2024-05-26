package person

import (
	"testing"
)

// TestPersonString tests the String method of the Person struct.
// It verifies that the method produces the correct string representation
// based on various attributes of the Person struct.
func TestPersonString(t *testing.T) {
	tests := []struct {
		person   Person
		expected string
	}{
		// Test case: Alice has a negative HasToPay indicating she will collect money
		{
			Person{Name: "Alice", InitialContribution: 50, HasToPay: -20, CanAfford: 100},
			"Alice -> -20.00 (Will collect)",
		},
		// Test case: Bob has a positive HasToPay and zero CanAfford indicating maximum affordability
		{
			Person{Name: "Bob", InitialContribution: 30, HasToPay: 20, CanAfford: 0},
			"Bob -> 20.00 (Max Afford)",
		},
		// Test case: Charlie has a positive HasToPay and sufficient affordability
		{
			Person{Name: "Charlie", InitialContribution: 10, HasToPay: 40, CanAfford: 100},
			"Charlie -> 40.00",
		},
		// Test case: Diana has zero HasToPay and zero CanAfford indicating maximum affordability
		{
			Person{Name: "Diana", InitialContribution: 20, HasToPay: 0, CanAfford: 0},
			"Diana -> 0.00 (Max Afford)",
		},
		// Test case: Eve has a positive HasToPay and sufficient affordability
		{
			Person{Name: "Eve", InitialContribution: 0, HasToPay: 50, CanAfford: 100},
			"Eve -> 50.00",
		},
	}

	// Iterate over each test case and verify the output of the String method
	for _, test := range tests {
		result := test.person.String()
		if result != test.expected {
			t.Errorf("Expected %s but got %s", test.expected, result)
		}
	}
}
