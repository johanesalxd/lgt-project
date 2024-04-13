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
		{description: "convert 3 to III", input: 3, want: "III"},
		{description: "convert 4 to IV", input: 4, want: "IV"},
		{description: "convert 5 to V", input: 5, want: "V"},
		{description: "convert 1984 to MCMLXXXIV", input: 1984, want: "MCMLXXXIV"},
		{description: "convert 3999 to MMMCMXCIX", input: 3999, want: "MMMCMXCIX"},
		{description: "convert 2014 to MMXIV", input: 2014, want: "MMXIV"},
		{description: "convert 1006 to MVI", input: 1006, want: "MVI"},
		{description: "convert 798 to DCCXCVIII", input: 798, want: "DCCXCVIII"},
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
