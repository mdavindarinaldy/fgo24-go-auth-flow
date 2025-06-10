package utils

import "fmt"

func GetInputInt(prompt string) (int, error) {
	fmt.Print(prompt)
	var input int
	_, err := fmt.Scanln(&input)
	return input, err
}

func GetInputString(prompt string) (string, error) {
	fmt.Print(prompt)
	var input string
	_, err := fmt.Scanln(&input)
	return input, err
}
