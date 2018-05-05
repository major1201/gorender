package goutils

import "strconv"

// ToInt safely parse a string to a integer with a default value 0 when parse fails
func ToInt(s string) int {
	return ToIntDv(s, 0)
}

// ToIntDv safely parse a string to a integer with a specified default value
func ToIntDv(s string, dv int) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		return dv
	}
	return i
}
