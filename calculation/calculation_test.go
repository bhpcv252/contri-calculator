package calculation

import (
	"testing"

	"github.com/bhpcv252/contri-calculator/person"
)

func TestCalculateContribution(t *testing.T) {
	group := []person.Person{
		{Name: "John", Amount: 20},
		{Name: "John", Amount: 20},
		{Name: "John", Amount: 20},
	}
	finalAmount := float32(100)

	expected := float32(float32(100) / 3)

	result := CalculateContribution(group, finalAmount)
	if result != expected {
		t.Errorf("Expected each contribution to be %.2f but got %.2f", expected, result)
	}

}

func TestCalculateEachContribution(t *testing.T) {
	group := []person.Person{
		{Name: "John", Amount: 20},
		{Name: "John", Amount: 30},
		{Name: "John", Amount: 50},
	}
	eachContribution := float32(40)

	expectedAmounts := []float32{20, 10, -10}

	CalculateEachContribution(group, eachContribution)

	for i, person := range group {
		if person.Amount != expectedAmounts[i] {
			t.Errorf("Expected amount for %s to be %.2f but got %.2f", person.Name, expectedAmounts[i], person.Amount)
		}
	}
}
