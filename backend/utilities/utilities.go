package utilities

// пакет utilities содержит вспомогательные функции для работы с Artist ID в artistHandler.go

import (
	"groupie-tracker/backend/data"
	"unicode"
)

// функция IsValid проверяет, является ли id допустимым. Она возвращает true, если строка не пустая и состоит только из цифр, иначе возвращает false.
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

// StartsWithZero: Эта функция проверяет, начинается ли строка идентификатора с нуля. Она возвращает true, если строка не равна "0" и первый символ равен '0', иначе возвращает false. Это нужно чтобы не срабатывали URL типа /artist/?id=000004
func StartsWithZero(id string) bool {
	if id != "0" && id[0] == '0' {
		return true
	}
	return false
}

// IsInRange: Эта функция проверяет, находится id в допустимом диапазоне. Она принимает целое число и возвращает true, если оно больше 0 и меньше или равно количеству артистов в data.Artists, иначе возвращает false.
func IsInRange(id int) bool {
	if id < 1 || id > len(data.Artists) {
		return false
	}
	return true
}
