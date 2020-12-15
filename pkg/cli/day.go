package cli

import (
	"fmt"
	"strconv"
	"strings"
)

type Day interface {
	Run(parts Parts) error
	Help() []string
}

type Days []Day

func (days Days) Run(args []string) error {
	if isHelpArgs(args) {
		return showHelp()
	}
	if len(args) == 2 && args[0] == "help" {
		dayIndex, err := days.getDayIndex(args[1])
		if err != nil {
			return err
		}
		fmt.Println(strings.Join(days[dayIndex].Help(), "\n"))
		fmt.Println()
		return nil
	}
	if len(args) == 1 || len(args) == 2 {
		dayIndex, err := days.getDayIndex(args[0])
		if err != nil {
			return err
		}
		parts := bothParts
		if len(args) == 2 {
			p, err := strconv.Atoi(args[1])
			if err != nil {
				return err
			}
			if p < 1 || p > 2 {
				return fmt.Errorf("<part> must be 1 or 2")
			}
			if p == 1 {
				parts = part1
			}
			if p == 2 {
				parts = part2
			}
		}
		return days[dayIndex].Run(parts)
	}
	return nil
}

func (days Days) getDayIndex(input string) (int, error) {
	day, err := parseDay(input)
	if err != nil {
		return 0, err
	}
	if day > len(days) {
		return 0, fmt.Errorf("day %d is not implemented yet", day)
	}
	return day - 1, nil
}

func showHelp() error {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("  <binary> [--help]")
	fmt.Println("    Show this help.")
	fmt.Println()
	fmt.Println("  <binary> help <day>")
	fmt.Println("    Show information about a specific day.")
	fmt.Println()
	fmt.Println("  <binary> <day> [<part>]")
	fmt.Println("    Runs a specific day. Arguments:")
	fmt.Println("    - <day>: A number between 1 and 25 (inclusive)")
	fmt.Println("    - <part>: 1 or 2. If missing, both parts are run.")
	return nil
}

func isHelpArgs(args []string) bool {
	return len(args) == 0 || (len(args) == 0 && args[0] == "--help")
}

func parseDay(input string) (int, error) {
	day, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	if day < 1 || day > 25 {
		return 0, fmt.Errorf("day must be between 1 and 25 (inclusive), not %d", day)
	}
	return day, nil
}

type Parts interface {
	Part1() bool
	Part2() bool
}

type parts int

func (p parts) Part1() bool {
	return p&1 == 1
}

func (p parts) Part2() bool {
	return p&2 == 2
}

const (
	part1     parts = 1
	part2     parts = 2
	bothParts parts = 3
)
