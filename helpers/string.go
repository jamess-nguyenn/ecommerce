package helpers

import (
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func generateStringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

func GenerateString(length int) string {
	return generateStringWithCharset(length, charset)
}

func GenerateLetter(length int) string {
	return generateStringWithCharset(length, letters)
}

func GenerateUUID() string {
	return uuid.NewString()
}

func ToInt(s string) int {
	number, _ := strconv.Atoi(s)
	return number
}

func ToUint64(s string) uint64 {
	number, _ := strconv.ParseUint(s, 10, 64)
	return number
}

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}
