package scanning

import (
	"regexp"
	"strconv"
)

func toNumber(s string) (match string, ok bool) {
	re := regexp.MustCompile(`^\d+(\.\d+)?`)
	matches := re.FindStringSubmatch(s)

	if len(matches) == 0 {
		return "", false
	}

	return matches[0], true
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
