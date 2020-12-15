package day2

import (
	"fmt"
	"io/ioutil"
	"strings"

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
	rawInput, err := ioutil.ReadFile(d.inputFilepath)
	if err != nil {
		return err
	}
	rawLines := strings.Split(string(rawInput), "\n")
	pwRecs := make([]*passwordRecord, len(rawLines)-1)
	for i := range pwRecs {
		pwRecs[i], err = passwordRecordFromLine(rawLines[i])
		if err != nil {
			return fmt.Errorf("error in line %d: %+v", i+1, err)
		}
	}
	d.pwRecords = pwRecs
	return nil
}

func (d *sday) Part1() error {
	validPasswordCount := 0
	for _, pwRec := range d.pwRecords {
		if pwRec.isValid() {
			validPasswordCount++
		}
	}
	fmt.Printf("There are %d valid passwords.\n", validPasswordCount)
	return nil
}

func (d *sday) Part2() error {
	return nil
}
