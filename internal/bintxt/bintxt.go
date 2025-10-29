// Package bintxt
package bintxt

import (
	"fmt"
	"strconv"
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

func Decode(s string) (string, error) {
	if s == "" {
		return "", nil
	}
	numbers := strings.Split(s, " ")
	b := new(strings.Builder)
	for i, num := range numbers {
		base, bitsize := 2, 8
		uint64Val, err := strconv.ParseUint(num, base, bitsize)
		if err != nil {
			return "", &ParseError{err: err.(*strconv.NumError), Seg: i + 1}
		}
		// Return error if number of bits is not exactly 8
		if len(num) != 8 {
			return "", fmt.Errorf("segment %d: parsing: %q: syntax: binary value must be exactly 8 bits, was %d bits", i+1, num, len(num))
		}
		// Safe because bitsize 8 is passed to PrseUint
		byteVal := byte(uint64Val)
		b.WriteByte(byteVal)
	}
	return b.String(), nil
}

type ParseError struct {
	err *strconv.NumError
	Seg int
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("segment %d: parsing: %q: %v", e.Seg, e.err.Num, e.err.Err.Error())
}

func (e *ParseError) Unwrap() error { return e.err }
