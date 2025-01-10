package utils

import (
	"regexp"
)

func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	compile := regexp.MustCompile(pattern)
	return compile.MatchString(email)
}

func VerifyPasswordFormat(password string) bool {
	pattern := `^(?![0-9]+$)(?![a-zA-Z]+$)[0-9a-zA-Z]{8, 16}$`
	compile := regexp.MustCompile(pattern)
	return compile.MatchString(password)
}
