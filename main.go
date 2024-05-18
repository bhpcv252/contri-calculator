package main

import (
	"fmt"

	"github.com/bhpcv252/contri-calculator/calculation"
	"github.com/bhpcv252/contri-calculator/person"
)

func main() {
	fmt.Println("\n-----------------")
	peopleCount := getIntInput("How many people are in the group? : ")

	group := make([]person.Person, 0, peopleCount)

	fmt.Println("\n-----------------")
	fmt.Println("Enter people's name and initial contribution: (example: John 20) ")

	for i := 0; i < peopleCount; i++ {
		name, amount := getPersonInput("(Name Amount): ")
		group = append(group, person.Person{Name: name, Amount: amount})
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
