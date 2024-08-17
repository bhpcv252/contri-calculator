package main

import (
	"fmt"
	"log"

	"github.com/bhpcv252/contri-calculator/calculation"
	"github.com/bhpcv252/contri-calculator/person"
)

func main() {

	finalAmount := float32(0)

	group := make([]*person.Person, 0)

	fmt.Println("\n-----------------")
	fmt.Println("Enter person's name, amount paid (optional) and how much they can afford (optional): (example: John 20 50). Leave empty to finish.")

	for {
		name, amount, afford, err := getPersonInput("(Name Paid Budget): ")
		if err != nil {
			break
		}
		finalAmount += amount
		group = append(group, &person.Person{Name: name, InitialContribution: amount, AffordLimit: afford, CanAfford: afford})
	}

	if len(group) == 0 {
		log.Fatal("No people were entered.")
	}

	fmt.Println("\n-----------------")
	fmt.Printf("Final Amount: %.2f\n", finalAmount)

	calculation.CalculateContributions(group, finalAmount, 0)

	fmt.Println("\n-----------------")
	for _, p := range group {
		fmt.Println(p)
	}

}
