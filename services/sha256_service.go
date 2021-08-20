package services

import (
	"crypto/sha256"
	"fmt"
)

func SHA256Encoder(uncrip string) string {
	aux := sha256.Sum256([]byte(uncrip))
	return fmt.Sprintf("%x", aux)
}
