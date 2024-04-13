package property

import "testing"

var cases = []struct {
	description string
	arabic      uint16
	roman       string
}{
	{description: "convert 1 to I (vice versa)", arabic: 1, roman: "I"},
	{description: "convert 2 to II (vice versa)", arabic: 2, roman: "II"},
	{description: "convert 3 to III (vice versa)", arabic: 3, roman: "III"},
	{description: "convert 4 to IV (vice versa)", arabic: 4, roman: "IV"},
	{description: "convert 5 to V (vice versa)", arabic: 5, roman: "V"},
	{description: "convert 1984 to MCMLXXXIV (vice versa)", arabic: 1984, roman: "MCMLXXXIV"},
	{description: "convert 3999 to MMMCMXCIX (vice versa)", arabic: 3999, roman: "MMMCMXCIX"},
	{description: "convert 2014 to MMXIV (vice versa)", arabic: 2014, roman: "MMXIV"},
	{description: "convert 1006 to MVI (vice versa)", arabic: 1006, roman: "MVI"},
	{description: "convert 798 to DCCXCVIII (vice versa)", arabic: 798, roman: "DCCXCVIII"},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(test.description, func(t *testing.T) {
			got := ConvertToRoman(test.arabic)

			if got != test.roman {
				t.Errorf("got %q want %q", got, test.roman)
			}
		})
	}
}

func TestArabicNumerals(t *testing.T) {
	for _, test := range cases {
		t.Run(test.description, func(t *testing.T) {
			got := ConvertToArabic(test.roman)

			if got != test.arabic {
				t.Errorf("got %q want %q", got, test.arabic)
			}
		})
	}
}
