package scanning

import "regexp"

func isNumber(s string) bool {
	re := regexp.MustCompile(`^\d+((\.\d+)+)?.?$`)
	return re.MatchString(s)
}
