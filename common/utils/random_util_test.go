package utils

import (
	"fmt"
	"net/url"
	"testing"
)

func TestName(t *testing.T) {
	escape := url.QueryEscape("#gtr32.jpeg")
	fmt.Println(escape)
}
