package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandInt generates a random value between the provided min and max values
func RandInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandString generates a random alphabetical string of length n
func RandString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandOwner generates a random owner string
func RandomOwner() string {
	return RandString(6)
}

func RandomMoney() int64 {
	return RandInt(0, 1000000)
}

func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD, KRW}
	n := len(currencies)

	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@mmail.com", RandString(6))
}
