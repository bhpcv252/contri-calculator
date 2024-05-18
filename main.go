package main

import (
	"fmt"
	"log"

	"github.com/bhpcv252/contri-calculator/calculation"
	"github.com/bhpcv252/contri-calculator/person"
)

func main() {

	group := make([]person.Person, 0)

	fmt.Println("\n-----------------")
	fmt.Println("Enter people's name and initial contribution: (example: John 20). Leave empty to finish.")

	for {
		name, amount, err := getPersonInput("(Name Amount): ")
		if err != nil {
			break
		}
		group = append(group, person.Person{Name: name, Amount: amount})
	}

	if len(group) == 0 {
		log.Fatal("No people were entered.")
	}

	fmt.Println("\n-----------------")
	finalAmount := getFloatInput("Enter the final amount which was spent? : ")

	eachContribution := calculation.CalculateContribution(group, finalAmount)

	fmt.Println("\n-----------------")
	fmt.Printf("Each contribution is: %.2f\n", eachContribution)

	calculation.CalculateEachContribution(group, eachContribution)

	fmt.Println("\n-----------------")
	calculation.PrintGroupStatus(group)

}
