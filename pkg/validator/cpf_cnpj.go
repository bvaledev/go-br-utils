package validator

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func IsCPFCNPJ(cpfcnpj string) bool {
	value := removeNonDigits(cpfcnpj)
	if value == "" {
		return false
	}
	if hasValidLength(value, true) {
		return IsCNPJ(value)
	} else if hasValidLength(value, false) {
		return IsCPF(value)
	}
	return false
}

func IsCPF(cpf string) bool {
	value := removeNonDigits(cpf)
	if !hasValidLength(value, false) || hasTheSameDigit(value) {
		return false
	}
	firstDigit := calculateCPFDigit(value, 10)
	secondDigit := calculateCPFDigit(value, 11)
	digits := fmt.Sprintf("%d%d", firstDigit, secondDigit)
	return extractCheckDigit(value) == digits
}

func IsCNPJ(cnpj string) bool {
	value := removeNonDigits(cnpj)
	if !hasValidLength(value, true) || hasTheSameDigit(value) {
		return false
	}
	firstDigit := calculateCNPJDigit(value, 12)
	secondDigit := calculateCNPJDigit(value, 13)
	digits := fmt.Sprintf("%d%d", firstDigit, secondDigit)
	return extractCheckDigit(value) == digits
}

func removeNonDigits(value string) string {
	regex := regexp.MustCompile(`\D+`)
	return regex.ReplaceAllString(value, "")
}

func hasTheSameDigit(value string) bool {
	digits := strings.Split(value, "")
	firstDigit := digits[0]
	equalLengthQuantity := 0
	for _, digit := range digits {
		if firstDigit == digit {
			equalLengthQuantity += 1
		}
	}
	return equalLengthQuantity == len(value)
}

func hasValidLength(value string, IsCNPJ bool) bool {
	length := len(value)
	if IsCNPJ {
		return length == 14
	} else {
		return length == 11
	}
}

func extractCheckDigit(value string) string {
	return value[len(value)-2:]
}

func calculateCPFDigit(value string, factor int) int {
	total := 0
	for _, digit := range strings.Split(value, "") {
		if factor > 1 {
			digitNumber, _ := strconv.Atoi(digit)
			total += digitNumber * factor
			factor--
		}
	}
	rest := total % 11
	if rest < 2 {
		return 0
	}
	return 11 - rest
}

func calculateCNPJDigit(value string, factor int) int {
	multiplier1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	multiplier2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	var multiplier []int = multiplier1
	if factor > 12 {
		multiplier = multiplier2
	}
	total := 0
	for i, digit := range strings.Split(value[:factor], "") {
		digitNumber, _ := strconv.Atoi(string(digit))
		total += digitNumber * multiplier[i]
	}
	rest := total % 11
	if rest < 2 {
		return 0
	}
	return 11 - rest
}
