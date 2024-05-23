package calculation

import (
	"fmt"

	"github.com/bhpcv252/contri-calculator/person"
)

func AffordCount(group []person.Person) int {
	totalMemebers := 0

	for _, person := range group {
		if person.HasToPay+person.InitialContribution == person.MaxAfford {
			continue
		}
		totalMemebers++
	}

	return totalMemebers
}

func CalculateContribution(group []person.Person, finalAmount float32) float32 {
	if len(group) == 0 {
		return 0
	}

	totalMemebers := AffordCount(group)

	return finalAmount / float32(totalMemebers)
}

func CalculateEachContribution(group []person.Person, finalAmount float32, includeInitial bool) {
	if finalAmount <= 0 || AffordCount(group) == 0 {
		return
	}

	eachPay := CalculateContribution(group, finalAmount)

	for i := range group {

		if group[i].HasToPay+group[i].InitialContribution == group[i].MaxAfford {
			continue
		}

		amountToDeduct := float32(0)
		if group[i].HasToPay+eachPay > group[i].MaxAfford {
			if includeInitial {
				amountToDeduct = group[i].MaxAfford - group[i].InitialContribution
			} else {
				amountToDeduct = group[i].MaxAfford
			}
			group[i].HasToPay += amountToDeduct
		} else {
			if includeInitial {
				amountToDeduct = eachPay - group[i].InitialContribution
			} else {
				amountToDeduct = eachPay
			}
			group[i].HasToPay += amountToDeduct
		}

		finalAmount -= amountToDeduct + group[i].InitialContribution
	}

	CalculateEachContribution(group, finalAmount, false)
}

func PrintGroupStatus(group []person.Person) {
	for _, person := range group {
		fmt.Println(person)
	}
}
