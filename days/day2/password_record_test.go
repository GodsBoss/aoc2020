package day2

import (
	"testing"
)

func TestPasswordRecordFromLine(t *testing.T) {
	line := "9-14 t: tptjdnnzkvjqbthm"
	pwRec, err := passwordRecordFromLine(line)
	if err != nil {
		t.Fatalf("failed getting password record from file: %+v", err)
	}
	if pwRec == nil {
		t.Error("expected a password record")
	}
}

func TestPasswordRecordFromLineFails(t *testing.T) {
	testcases := map[string]string{
		"broken line":    "Hello, world!",
		"too many lines": "9-14 t: tptjdnnzkvjqbthm   9-14 t: tptjdnnzkvjqbthm   9-14 t: tptjdnnzkvjqbthm",
		"broken minimum": "9999999999999999999999-14 t: tptjdnnzkvjqbthm",
		"broken maximum": "9-7777777777777777777777 t: tptjdnnzkvjqbthm",
		"min > max":      "14-9 t: tptjdnnzkvjqbthm",
	}

	for name, line := range testcases {
		t.Run(
			name,
			func(t *testing.T) {
				pwRec, err := passwordRecordFromLine(line)
				if err == nil {
					t.Error("expected error")
				}
				if pwRec != nil {
					t.Errorf("expected no password record, got %+v", pwRec)
				}
			},
		)
	}
}
