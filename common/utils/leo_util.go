package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"regexp"
)

// IsValidLeoAddress validate hex address
func IsValidLeoAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^aleo[0-9a-fA-F]{59}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}
