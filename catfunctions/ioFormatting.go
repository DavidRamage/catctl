package catfunctions

import (
	"fmt"
	"strconv"
	"unicode"
)

func ToHz(s string) (string, error) {
	var multiplier float64
	unit := s[len(s)-1]
	//input is already in Hz
	if unicode.IsDigit(rune(unit)) {
		hz, err := strconv.Atoi(s)
		if err != nil {
			return "0", fmt.Errorf("invalid frequency format: %s", s)
		}
		return fmt.Sprintf("%08d", hz), nil
	}
	freq, err := strconv.ParseFloat(s[:len(s)-2], 64)
	if err != nil {
		return "0", fmt.Errorf("invalid frequency format: %s", s)
	}
	switch unit {
	case 'k', 'K':
		multiplier = 1000
	case 'm', 'M':
		multiplier = 1000000
	default:
		return "0", fmt.Errorf("invalid frequency format: %s", s)
	}
	//hz := strconv.FormatInt(int64(freq*float64(multiplier)), 10)
	hz := freq * multiplier
	return fmt.Sprintf("%08d", hz), nil
}
