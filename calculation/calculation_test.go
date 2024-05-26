package calculation

import (
	"testing"

	"github.com/bhpcv252/contri-calculator/person"
)

// TestAffordCount tests the AffordCount function to ensure it correctly counts
// the number of people in a group who can afford to contribute.
func TestCalculateContributions(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name             string
		group            []*person.Person
		finalAmount      float32
		expectedPayments []float32
	}{
		{
			name: "Equal InitialContribution",
			group: []*person.Person{
				{InitialContribution: 0, CanAfford: -1, HasToPay: 0},
				{InitialContribution: 0, CanAfford: -1, HasToPay: 0},
				{InitialContribution: 1000, CanAfford: 100, HasToPay: 0},
				{InitialContribution: 0, CanAfford: -1, HasToPay: 0},
			},
			finalAmount:      1000,
			expectedPayments: []float32{300, 300, -900, 300},
		},
		{
			name: "Mixed InitialContribution and affordability",
			group: []*person.Person{
				{InitialContribution: 10, CanAfford: 50, HasToPay: 0},
				{InitialContribution: 0, CanAfford: -1, HasToPay: 0},
				{InitialContribution: 1000, CanAfford: 100, HasToPay: 0},
				{InitialContribution: 100, CanAfford: -1, HasToPay: 0},
			},
			finalAmount:      400,
			expectedPayments: []float32{40, 125, -900, 25},
		},
		{
			name: "Maximum affordability constraints",
			group: []*person.Person{
				{InitialContribution: 100, CanAfford: 200, HasToPay: 0},
				{InitialContribution: 200, CanAfford: 150, HasToPay: 0},
				{InitialContribution: 300, CanAfford: 100, HasToPay: 0},
				{InitialContribution: 400, CanAfford: 50, HasToPay: 0},
			},
			finalAmount:      2000,
			expectedPayments: []float32{100, -50, -200, -350},
		},
		{
			name: "Mixed complicated affordability constraints",
			group: []*person.Person{
				{InitialContribution: 10, CanAfford: 50, HasToPay: 0},
				{InitialContribution: 0, CanAfford: -1, HasToPay: 0},
				{InitialContribution: 1000, CanAfford: 110, HasToPay: 0},
				{InitialContribution: 90, CanAfford: -1, HasToPay: 0},
			},
			finalAmount:      400,
			expectedPayments: []float32{40, 120, -890, 30},
		},
	}

	// Loop through test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			CalculateContributions(tc.group, tc.finalAmount, 0)
			for i, person := range tc.group {
				if person.HasToPay != tc.expectedPayments[i] {
					t.Errorf("%s failed: Expected HasToPay %.2f but got %.2f", tc.name, tc.expectedPayments[i], person.HasToPay)
				}
			}
		})
	}
}
