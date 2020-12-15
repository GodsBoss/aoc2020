package day1

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func readInput(filename string) ([]int, error) {
	rawInput, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	rawLines := strings.Split(string(rawInput), "\n")
	lines := make([]int, len(rawLines)-1)
	for i := range lines {
		line, err := strconv.Atoi(rawLines[i])
		if err != nil {
			return nil, fmt.Errorf("error in line %d: %+v", i+1, err)
		}
		lines[i] = line
	}
	return lines, nil
}
