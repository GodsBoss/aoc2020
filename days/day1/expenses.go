package day1

import (
	"github.com/GodsBoss/aoc2020/pkg/cli"

	"fmt"
)

func NewDay(inputFilepath string) cli.Day {
	return cli.SimpleDay(
		&sday{
			inputFilepath: inputFilepath,
		},
		[]string{
			"Day 1: Report repair",
		},
	)
}

type sday struct {
	inputFilepath string
	lines         []int
}

func (d *sday) Init() error {
	lines, err := readInput(d.inputFilepath)
	if err != nil {
		return err
	}
	d.lines = lines
	return nil
}

func (d *sday) Part1() error {
	firstPartSolution, err := findCombination(d.lines, 2020, 2)
	if err != nil {
		return err
	}
	fmt.Printf("Solution the first part is %d.\n", multiply(firstPartSolution))
	return nil
}

func (d *sday) Part2() error {
	secondPartSolution, err := findCombination(d.lines, 2020, 3)
	if err != nil {
		return err
	}
	fmt.Printf("Solution to the second part is %d.\n", multiply(secondPartSolution))
	return nil
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
