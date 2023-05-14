package internal

import (
	"math/rand"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateSalt() string {
	salt := strings.Builder{}
	salt.Grow(6)
	for i := 0; i < 6; i++ {
		//51 is the size of charset
		salt.WriteByte(charset[rand.Intn(51)])
	}
	return salt.String()
}
