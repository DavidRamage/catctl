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
	hz := freq * multiplier
	return fmt.Sprintf("%08d", int(hz)), nil
}

func FmtFrequencyOut(rawFreq string, unit string, prefixLen int) (string, error) {
	frequency, err := strconv.Atoi(rawFreq[prefixLen:])
	if err != nil {
		return "Error", fmt.Errorf("unable to process frequency string: %s", rawFreq)
	}
	if unit == "hz" {
		return fmt.Sprintf("%d hz", frequency), nil
	} else if unit == "khz" {
		return fmt.Sprintf("%d khz", frequency/1000), nil
	} else if unit == "mhz" {
		return fmt.Sprintf("%d mhz", frequency/1000000), nil
	} else {
		return "Error", fmt.Errorf("Invalid frequency unit: %s", unit)
	}
}
