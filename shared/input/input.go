package input

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SplitIntoLines(s string) []string {
	return strings.Split(strings.TrimSpace(s), "\n")
}

func Parse(re *regexp.Regexp, s string, values ...any) error {
	m := re.FindStringSubmatch(s)
	if m == nil {
		return fmt.Errorf("Parse: line doesn't match regexp: %#v", s)
	}
	captures := m[1:]

	if len(values) != len(captures) {
		return fmt.Errorf("Parse: wrong number of values passed in: %d values for %d captures", len(values), len(captures))
	}

	for i, val := range values {
		capture := captures[i]

		switch val := val.(type) {
		case *string:
			*val = capture

		case *int:
			n, err := strconv.Atoi(capture)
			if err != nil {
				return fmt.Errorf("Parse: capture is not an int: %#v", capture)
			}
			*val = n

		case *float64:
			f, err := strconv.ParseFloat(capture, 64)
			if err != nil {
				return fmt.Errorf("Parse: capture is not a float: %#v", capture)
			}
			*val = f

		default:
			return fmt.Errorf("Parse: don't know how to convert to type %T", val)
		}
	}

	return nil
}
