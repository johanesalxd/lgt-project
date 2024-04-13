package property

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		description string
		input       int
		want        string
	}{
		{description: "convert 1 to I", input: 1, want: "I"},
		{description: "convert 2 to II", input: 2, want: "II"},
	}

	for _, test := range cases {
		t.Run(test.description, func(t *testing.T) {
			got := ConvertToRoman(test.input)

			if got != test.want {
				t.Errorf("got %q want %q", got, test.want)
			}
		})
	}
}
