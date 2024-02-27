package helper

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSHA256(input string) string {
	hash := sha256.New()
	hash.Write([]byte(input))
	hashedBytes := hash.Sum(nil)
	hashedString := hex.EncodeToString(hashedBytes)

	return hashedString
}
