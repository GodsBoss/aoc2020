package day2

import (
	"github.com/GodsBoss/aoc2020/pkg/cli"
)

type sday struct {
	inputFilepath string
	pwRecords     []*passwordRecord
}

func NewDay(inputFilepath string) cli.Day {
	return cli.SimpleDay(
		&sday{
			inputFilepath: inputFilepath,
		},
		[]string{
			"Day 2: Password Philosophy",
		},
	)
}

func (d *sday) Init() error {
	return nil
}

func (d *sday) Part1() error {
	return nil
}

func (d *sday) Part2() error {
	return nil
}
