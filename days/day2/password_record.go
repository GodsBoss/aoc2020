package day2

import (
	"fmt"
	"regexp"
	"strconv"
)

type passwordRecord struct {
	policy   policy
	password string
}

func passwordRecordFromLine(line string) (*passwordRecord, error) {
	parts := pwRecRE.FindAllStringSubmatch(line, -1)
	if len(parts) != 1 {
		return nil, fmt.Errorf("invalid line '%s'", line)
	}
	min, err := strconv.Atoi(parts[0][1])
	if err != nil {
		return nil, fmt.Errorf("invalid minimum: '%s'", parts[0][1])
	}
	max, err := strconv.Atoi(parts[0][2])
	if err != nil {
		return nil, fmt.Errorf("invalid maximum: '%s'", parts[0][2])
	}
	if min > max {
		return nil, fmt.Errorf("minimum %d is greater than maximum %d", min, max)
	}
	ch := parts[0][3]
	pw := parts[0][4]

	return &passwordRecord{
		policy: policy{
			min: min,
			max: max,
			ch:  ch,
		},
		password: pw,
	}, nil
}

var pwRecRE = regexp.MustCompile("([0-9]+)-([0-9]+) ([a-z]): ([a-z]+)")

type policy struct {
	min int
	max int
	ch  string
}
