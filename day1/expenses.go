package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
}

func run() error {
	rawInput, err := ioutil.ReadFile("input.txt")
	if err != nil {
		return err
	}
	rawLines := strings.Split(string(rawInput), "\n")
	lines := make([]int, len(rawLines)-1)
	for i := range lines {
		line, err := strconv.Atoi(rawLines[i])
		if err != nil {
			fmt.Printf("Error in line %d: %+v\n", i+1, err)
		}
		lines[i] = line
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
