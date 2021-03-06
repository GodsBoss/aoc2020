package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func (p policy) matches(pw string) bool {
	cnt := strings.Count(pw, p.ch)
	return (p.min <= cnt) && (cnt <= p.max)
}

func (p policy) altermatches(pw string) bool {
	matchingPositions := 0
	positions := []int{p.min - 1, p.max - 1}
	for _, pos := range positions {
		if pos < len(pw) && string(pw[pos]) == p.ch {
			matchingPositions++
		}
	}
	return matchingPositions == 1
}

func (rec *passwordRecord) isValid() bool {
	return rec.policy.matches(rec.password)
}

func (rec *passwordRecord) isReallyValid() bool {
	return rec.policy.altermatches(rec.password)
}
