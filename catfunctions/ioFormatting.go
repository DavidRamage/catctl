package catfunctions

import (
	"fmt"
	"strconv"
	"unicode"
)

func ToHz(s string) (string, error) {
	var multiplier int
	unit := s[len(s)-1]
	if unicode.IsDigit(rune(unit)) {
		return fmt.Sprintf("%09d", s), nil
	}
	freq, err := strconv.ParseFloat(s[:len(s)-1], 64)
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
	hz := strconv.FormatInt(int64(freq*float64(multiplier)), 10)
	return fmt.Sprintf("%09d", hz), nil
}
