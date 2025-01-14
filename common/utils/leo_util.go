package utils

import (
	"regexp"
)

// IsValidLeoAddress validate hex address
func IsValidLeoAddress(iaddress string) bool {
	re := regexp.MustCompile("^aleo[0-9a-zA-Z]{59}$")
	return re.MatchString(iaddress)
}
