package calculation

import (
	"testing"

	"github.com/bhpcv252/contri-calculator/person"
)

func TestAffordCount(t *testing.T) {
	group := []person.Person{
		{Name: "Sonu", InitialContribution: 20, HasToPay: 0, MaxAfford: 50},
		{Name: "Bob", InitialContribution: 50, HasToPay: 0, MaxAfford: 50},
		{Name: "Diana", InitialContribution: 200, HasToPay: 0, MaxAfford: 1000},
		{Name: "Charlie", InitialContribution: 100, HasToPay: 0, MaxAfford: 150},
	}
	expected := 3
	result := AffordCount(group)
	if result != expected {
		t.Errorf("Expected afford count to be %d but got %d", expected, result)
	}
}

func TestCalculateContribution(t *testing.T) {
	group := []person.Person{
		{Name: "Alice", InitialContribution: 20},
		{Name: "Bob", InitialContribution: 50},
		{Name: "Diana", InitialContribution: 200},
		{Name: "Charlie", InitialContribution: 100},
	}
	finalAmount := float32(1000)

	expected := float32(float32(1000) / 4)

	result := CalculateContribution(group, finalAmount)
	if result != expected {
		t.Errorf("Expected each contribution to be %.2f but got %.2f", expected, result)
	}

}

func TestCalculateEachContribution(t *testing.T) {
	finalAmount := float32(1000)

	group := []person.Person{
		{Name: "Alice", InitialContribution: 20, MaxAfford: 50},
		{Name: "Bob", InitialContribution: 50, MaxAfford: 200},
		{Name: "Diana", InitialContribution: 200, MaxAfford: finalAmount},
		{Name: "Charlie", InitialContribution: 100, MaxAfford: 150},
	}

	expectedAmounts := []float32{30, 150, 400, 50}

	CalculateEachContribution(group, finalAmount, true)

	for i, person := range group {
		if person.HasToPay != expectedAmounts[i] {
			t.Errorf("Expected amount for %s to be %.2f but got %.2f", person.Name, expectedAmounts[i], person.HasToPay)
		}
	}
}

func TestCalculateEachContributionEdgeCases(t *testing.T) {
	// Case with no persons in group
	group := []person.Person{}
	finalAmount := float32(1000)
	expected := float32(0)
	result := CalculateContribution(group, finalAmount)
	if result != expected {
		t.Errorf("Expected each contribution to be %.2f but got %.2f", expected, result)
	}

	// Case with finalAmount <= 0
	group = []person.Person{
		{Name: "Alice", InitialContribution: 20, MaxAfford: 50},
		{Name: "Bob", InitialContribution: 50, MaxAfford: 200},
	}
	finalAmount = float32(0)
	expected = float32(0)
	CalculateEachContribution(group, finalAmount, true)
	for _, person := range group {
		if person.HasToPay != expected {
			t.Errorf("Expected each contribution to be %.2f but got %.2f", expected, person.HasToPay)
		}
	}

	// Case where no one can afford to pay
	group = []person.Person{
		{Name: "Alice", InitialContribution: 50, MaxAfford: 50},
		{Name: "Bob", InitialContribution: 200, MaxAfford: 200},
	}
	finalAmount = float32(1000)
	expected = float32(0)
	CalculateEachContribution(group, finalAmount, true)
	for _, person := range group {
		if person.HasToPay != expected {
			t.Errorf("Expected each contribution to be %.2f but got %.2f", expected, person.HasToPay)
		}
	}
}
