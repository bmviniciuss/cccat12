package testutils

import (
	"math/rand"
	"strconv"
	"time"
)

func GenerateRandomCPF() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// Generate 9 random digits
	digits := make([]int, 9)
	for i := range digits {
		digits[i] = r.Intn(10)
	}

	// First digit calculation
	d1 := calculateVerifierDigit(digits, 10)
	digits = append(digits, d1)
	d2 := calculateVerifierDigit(digits, 11)
	digits = append(digits, d2)

	// Convert digits to string
	cpf := ""
	for _, digit := range digits {
		cpf += strconv.Itoa(digit)
	}
	return cpf
}

func calculateVerifierDigit(digits []int, factor int) int {
	sum := 0
	for i := 0; i < len(digits); i++ {
		if factor < 2 {
			break
		}
		sum += digits[i] * factor
		factor--
	}
	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}
