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
		// Check if the input has valid number of parts
		if len(parts) < 2 || len(parts) > 3 {
			fmt.Println("Invalid input. Please Enter in the format 'Name Amount Afford'. Afford is optional")
			continue
		}
		// Extract name from the input
		name := parts[0]
		// Extract amount from the input and convert it to float32
		amount, err := strconv.ParseFloat(parts[1], 32)
		if err != nil {
			fmt.Println("Invalid input. Please Enter enter a valid number.")
			continue
		}

		// Initialize afford variable
		var afford float32
		// If affordability is provided in the input, extract it and convert it to float32
		if len(parts) == 3 {
			affordValue, err := strconv.ParseFloat(parts[2], 32)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number for afford.")
				continue
			}
			afford = float32(affordValue)
		} else {
			// If affordability is not provided, use the default value finalAmount
			afford = -1
		}

		// Return name, amount, and affordability
		return name, float32(amount), float32(afford), nil

	}
}
