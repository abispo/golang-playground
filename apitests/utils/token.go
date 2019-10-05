package utils

import (
	"crypto/rand"
	"fmt"
)

// TokenGenerator ...
func TokenGenerator() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
