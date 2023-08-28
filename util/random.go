package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// rand.Int63n(n) returns, as an int64, a non-negative pseudo-random number -> 0~n-1.
// 0+min <= n <= max-min+1-1+min -> min <= n <=max
// RandomInt generates a random integer between min and max.
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	alphabetLen := len(alphabet)

	for i := 0; i < n; i++ {
		// rand.Intn(n) returns, as an int, a non-negative pseudo-random number -> 0~n-1.
		randomIndex := rand.Intn(alphabetLen)
		// randomChar get a random character from alphabet by randomIndex.
		randomChar := alphabet[randomIndex]
		sb.WriteByte(randomChar)
	}
	return sb.String()
}

// RandomOwner generates a random owner name from alphabet of length 6
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	// rand.Intn(n) returns, as an int, a non-negative pseudo-random number -> 0~n-1.
	return currencies[rand.Intn(n)]
}
