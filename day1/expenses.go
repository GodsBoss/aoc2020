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

	for i := range lines {
		for j := range lines {
			if i == j {
				continue
			}
			if lines[i]+lines[j] == 2020 {
				fmt.Printf("%d (line %d) * %d line %d = %d\n", lines[i], i+1, lines[j], j+1, lines[i]*lines[j])
				return nil
			}
		}
	}

	return fmt.Errorf("no matching combination found")
}
