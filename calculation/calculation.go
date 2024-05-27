package calculation

import (
	"github.com/bhpcv252/contri-calculator/person"
)

// CalculateContributions calculates the exact amount each person needs to contribute.
func CalculateContributions(group []*person.Person, finalAmount float32, depth int) {
	// Limit recursion depth to avoid infinite recursion
	if depth > 10000 {
		return
	}

	// Base case: if final amount is zero or group is empty, stop the recursion
	if finalAmount <= 0 || len(group) == 0 {
		return
	}

	// Calculate the equal share for each person in the group
	equalShare := finalAmount / float32(len(group))

	// Initialize remaining amount and the next group iteration for further processing
	remaining := float32(0)
	nextGroupIteration := make([]*person.Person, 0)

	// Iterate over each person in the group
	for _, p := range group {
		// If the person's payment exceeds their affordability, adjust their payment
		if p.CanAfford != -1 && equalShare >= p.CanAfford {
			p.HasToPay += p.CanAfford
			if depth == 0 {
				p.HasToPay -= p.InitialContribution
			}
			// Calculate the remaining amount to be distributed
			remaining += equalShare - p.CanAfford

			// Person can't afford any more amount
			p.CanAfford = 0
		}

		// If the person can afford the equal share or has unlimited affordability
		if p.CanAfford == -1 || p.CanAfford > equalShare {
			p.HasToPay += equalShare
			if p.CanAfford != -1 {
				p.CanAfford -= equalShare
				if p.CanAfford < 0 {
					p.CanAfford = 0
				}
			}
			if depth == 0 {
				p.HasToPay -= p.InitialContribution
			}
			// Add person to next group for further processing
			nextGroupIteration = append(nextGroupIteration, p)
		}
	}

	// Recursively calculate the contribution for the remaining amount
	CalculateContributions(nextGroupIteration, remaining, depth+1)
}
