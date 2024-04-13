package property

import "strings"

type romanNumeral struct {
	value  uint16
	Symbol string
}

var allRomanNumerals = []romanNumeral{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(input uint16) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for input >= numeral.value {
			result.WriteString(numeral.Symbol)
			input -= numeral.value
		}
	}

	return result.String()
}

func ConvertToArabic(input string) uint16 {
	var result uint16 = 0

	for _, numeral := range allRomanNumerals {
		for strings.HasPrefix(input, numeral.Symbol) {
			result += numeral.value
			input = strings.TrimPrefix(input, numeral.Symbol)
		}
	}

	return result
}
