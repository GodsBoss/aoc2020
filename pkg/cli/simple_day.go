package cli

import (
	"fmt"
)

func SimpleDay(dayParts DayParts, help []string) Day {
	return &simpleday{
		parts: dayParts,
		help:  help,
	}
}

type DayParts interface {
	Init() error
	Part1() error
	Part2() error
}

type simpleday struct {
	parts DayParts
	help  []string
}

func (d *simpleday) Help() []string {
	return d.help
}

func (d *simpleday) Run(parts Parts) error {
	if err := d.parts.Init(); err != nil {
		return err
	}
	var errs []error
	if parts.Part1() {
		if err := d.parts.Part1(); err != nil {
			errs = append(errs, err)
		}
	}
	if parts.Part2() {
		if err := d.parts.Part2(); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) == 2 {
		return fmt.Errorf("errors running parts: %+v", errs)
	}
	if len(errs) == 1 {
		return errs[0]
	}
	return nil
}
