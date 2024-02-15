package utilities

import "unicode"

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

func IsRange(id int) bool {
	if id < 1 || id > 52 {
		return false
	}
	return true
}

func ContainsZero(id string) bool {
	if id[0] == '0' {
		return true
	}
	return false
}
