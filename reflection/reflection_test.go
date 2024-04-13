package reflection

import (
	"testing"
)

func TestWalk(t *testing.T) {
	var got []string

	want := "Chris"
	x := struct {
		name string
	}{want}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}
}
