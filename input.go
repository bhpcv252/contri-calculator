package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// getInput reads input from the user and returns it as a string.
func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')

	return strings.TrimSpace(input)
}

// getFloatInput prompts the user for input and returns it as a float32.
func getFloatInput(prompt string) float32 {
	for {
		// Get input from the user
		input := getInput(prompt)
		// Convert the input to a float32
		value, err := strconv.ParseFloat(input, 32)
		// If conversion is successful, return the value
		if err == nil {
			return float32(value)
		}
		// If conversion fails, prompt the user to enter a valid number
		fmt.Println("Invalid input. Please Enter a valid number.")
	}
}

// getPersonInput prompts the user to input person details and returns them.
func getPersonInput(prompt string) (string, float32, float32, error) {
	for {
		// Get input from the user
		input := getInput(prompt)
		// If input is empty, return an error indicating end of input
		if input == "" {
			return "", 0, 0, fmt.Errorf("end of input")
		}
		// Split the input by spaces
		parts := strings.Split(input, " ")
		// Check if the input has a valid number of parts
		if len(parts) < 1 || len(parts) > 3 {
			fmt.Println("Invalid input. Please enter in the format 'Name [Amount] [Afford]'. Amount and Afford are optional.")
			continue
		}
		// Extract name from the input
		name := parts[0]

		// Initialize amount and afford variables
		var amount float32
		var afford float32

		// If amount is provided in the input, extract it and convert it to float32
		if len(parts) >= 2 {
			amountValue, err := strconv.ParseFloat(parts[1], 32)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number for amount.")
				continue
			}
			amount = float32(amountValue)
		} else {
			// If amount is not provided, use the default value 0
			amount = 0
		}

		// If affordability is provided in the input, extract it and convert it to float32
		if len(parts) == 3 {
			affordValue, err := strconv.ParseFloat(parts[2], 32)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number for afford.")
				continue
			}
			afford = float32(affordValue)
		} else {
			// If affordability is not provided, use the default value -1
			afford = -1
		}

		// Return name, amount, and affordability
		return name, amount, afford, nil
	}
}
