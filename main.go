package main

import (
	"fmt"

	"log"
)

type Person struct {
	name   string
	amount float32
}

func (p Person) String() string {
	return fmt.Sprintf("%s - %f", p.name, p.amount)
}

func main() {
	group := make([]Person, 0, 3)
	input := Person{}

	fmt.Println("Enter people's name and initial contribution: (example: John 20) ")

	for i := 0; i < cap(group); i++ {
		_, err := fmt.Scanf("%s %f\n", &input.name, &input.amount)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Printf("Enter: %d %s", n, input)

		group = append(group, input)
	}

	fmt.Println(len(group), cap(group), group)

}
