package utilities

import (
	"groupie-tracker/backend/data"
	"unicode"
)

func IsValid(id string) bool {
	if id == "" {
		return false
	}
	for _, char := range id {
		if !unicode.IsDigit(char) {
			return false
		}
	}

	return true
}

func IsInRange(id int) bool {
	if id < 1 || id > len(data.Artists) {
		return false
	}
	return true
}

func StartsWithZero(id string) bool {
	if id[0] == '0' {
		return true
	}
	return false
}
