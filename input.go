package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

func getFloatInput(prompt string) float32 {
	for {
		input := getInput(prompt)
		value, err := strconv.ParseFloat(input, 32)
		if err == nil {
			return float32(value)
		}
		fmt.Println("Invalid input. Please Enter a valid number.")
	}
}

func getPersonInput(prompt string) (string, float32, error) {
	for {
		input := getInput(prompt)
		if input == "" {
			return "", 0, fmt.Errorf("end of input")
		}
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please Enter in the format 'Name Amount'.")
			continue
		}
		name := parts[0]
		amount, err := strconv.ParseFloat(parts[1], 32)
		if err != nil {
			fmt.Println("Invalid input. Please Enter enter a valid number.")
			continue
		}

		return name, float32(amount), nil

	}
}
