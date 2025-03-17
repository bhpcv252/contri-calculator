package main

import (
	"fmt"
	"log"
	"os"

	"github.com/bhpcv252/contri-calculator/calculation"
	"github.com/bhpcv252/contri-calculator/person"
)

var version = "v0.1.4" // current version

func main() {
	// Check if the first argument is either --version or -v
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Printf("Contri version: %s\n", version)
		return
	}

	finalAmount := float32(0)

	group := make([]*person.Person, 0)

	fmt.Println("\n-----------------")
	fmt.Print(
		"Enter person's name, amount paid (optional) and how much they can afford (optional).\nExample: John 20 50\nLeave empty to finish.\n\n",
	)

	for {
		name, amount, afford, err := getPersonInput("(Name Paid Budget): ")
		if err != nil {
			break
		}
		finalAmount += amount
		group = append(
			group,
			&person.Person{
				Name:                name,
				InitialContribution: amount,
				AffordLimit:         afford,
				CanAfford:           afford,
			},
		)
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
