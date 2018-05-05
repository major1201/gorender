package goutils

import (
	"github.com/google/uuid"
	"reflect"
	"regexp"
	"strings"
)

// EmptyStr is the empty string const
const EmptyStr = ""

// IsEmpty tells a string is empty or not
func IsEmpty(s string) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

// IsNotEmpty tells a string is not empty or is
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// IsBlank tells a string is blank or not, a string with all white blanks would be considered as true
func IsBlank(s string) bool {
	if len(s) == 0 {
		return true
	}
	reg := regexp.MustCompile(`^\s+$`)
	actual := reg.MatchString(s)
	if actual {
		return true
	}
	return false
}

// IsNotBlank tells a string is not blank or is
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// Trim cuts the blanks of a string, the beginning and the end
func Trim(str string) string {
	return strings.Trim(str, " ")
}

// TrimLeft cuts the left side blanks of a string
func TrimLeft(str string) string {
	return strings.TrimLeft(str, " ")
}

// TrimRight cuts the right side blanks of a string
func TrimRight(str string) string {
	return strings.TrimRight(str, " ")
}

// LeftPad pad a string with specified character to the left side
func LeftPad(s string, padStr string, length int) string {
	prefix := EmptyStr
	if len(s) < length {
		prefix = strings.Repeat(padStr, length-len(s))
	}
	return prefix + s
}

// RightPad pad a string with specified character to the right side
func RightPad(s string, padStr string, length int) string {
	postfix := EmptyStr
	if len(s) < length {
		postfix = strings.Repeat(padStr, length-len(s))
	}
	return s + postfix
}

// ZeroFill pad a string(usually a number string) with "0" to the left
func ZeroFill(s string, length int) string {
	const zeroStr = "0"
	return LeftPad(s, zeroStr, length)
}

// Len returns the length of a string using rune, it's useful when getting the length of a string including CJK characters
func Len(s string) int {
	return len([]rune(s))
}

// Index return the location of a string in another long string, if it doesn't exist, returns -1
// this function supports CJK characters
func Index(s, substr string) int {
	sRune := []rune(s)
	subRune := []rune(substr)
	if len(subRune) > len(sRune) {
		return -1
	}
	for i := 0; i < len(sRune)-len(subRune)+1; i++ {
		if reflect.DeepEqual(sRune[i:i+len(subRune)], subRune) {
			return i
		}
	}
	return -1
}

// Between return the middle part of a string from string "from" to string "to"
func Between(s, from, to string) string {
	indexFrom := Index(s, from)
	if indexFrom == -1 {
		return ""
	}
	runeS := []rune(s)
	indexTo := Index(string(runeS[indexFrom:]), to)
	if indexTo == -1 {
		return ""
	}
	return string(runeS[indexFrom+Len(from) : indexFrom+indexTo])
}

// UUID returns a random generated UUID string
func UUID() string {
	return strings.Replace(uuid.New().String(), "-", "", 4)
}
