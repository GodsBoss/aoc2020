package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	lines, err := readInput("input.txt")
	if err != nil {
		return err
	}

	firstPartSolution, err := findCombination(lines, 2020, 2)
	if err != nil {
		return err
	}
	fmt.Printf("Solution the first part is %d.\n", multiply(firstPartSolution))

	secondPartSolution, err := findCombination(lines, 2020, 3)
	if err != nil {
		return err
	}
	fmt.Printf("Solution to the second part is %d.\n", multiply(secondPartSolution))

	return nil
}

func multiply(factors []int) int {
	result := 1
	for i := range factors {
		result *= factors[i]
	}
	return result
}
