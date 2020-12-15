package main

import (
	"github.com/GodsBoss/aoc2020/pkg/cli"

	"fmt"
	"os"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	days := cli.Days{}
	return days.Run(args)
}
