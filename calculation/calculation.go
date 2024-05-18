package calculation

import (
	"fmt"

	"github.com/bhpcv252/contri-calculator/person"
)

func CalculateContribution(group []person.Person, finalAmount float32) float32 {
	return finalAmount / float32(len(group))
}

func CalculateEachContribution(group []person.Person, eachContribution float32) {
	for i := range group {
		group[i].Amount = eachContribution - group[i].Amount
	}
}

func PrintGroupStatus(group []person.Person) {
	for _, person := range group {
		fmt.Println(person)
	}
}
