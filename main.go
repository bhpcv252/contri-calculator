package main

import (
	"fmt"
	"log"

	"github.com/bhpcv252/contri-calculator/calculation"
	"github.com/bhpcv252/contri-calculator/person"
)

func main() {

	fmt.Println("\n-----------------")
	finalAmount := getFloatInput("Enter the final amount which was spent? : ")

	group := make([]person.Person, 0)

	fmt.Println("\n-----------------")
	fmt.Println("Enter people's name, initial contribution and how much they can afford (optional): (example: John 20 50). Leave empty to finish.")

	for {
		name, amount, afford, err := getPersonInput("(Name Amount Afford): ", finalAmount)
		if err != nil {
			break
		}
		group = append(group, person.Person{Name: name, InitialContribution: amount, MaxAfford: afford})
	}

	if len(group) == 0 {
		log.Fatal("No people were entered.")
	}

	eachContribution := calculation.CalculateContribution(group, finalAmount)

	fmt.Println("\n-----------------")
	fmt.Printf("Each contribution is: %.2f\n", eachContribution)

	calculation.CalculateEachContribution(group, finalAmount, true)

	fmt.Println("\n-----------------")
	calculation.PrintGroupStatus(group)

}
