package helper

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateOTPCode() (string, error) {
	maxInt := big.NewInt(1000000)
	n, err := rand.Int(rand.Reader, maxInt)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%06d", n.Int64()), nil
}
