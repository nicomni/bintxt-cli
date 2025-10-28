package bintxt

import (
	"fmt"
	"strings"
)

func Encode(s string) string {
	if len(s) == 0 {
		return ""
	}
	var parts []string
	for i := 0; i < len(s); i++ {
		parts = append(parts, fmt.Sprintf("%08b", s[i]))
	}
	return strings.Join(parts, " ")
}
