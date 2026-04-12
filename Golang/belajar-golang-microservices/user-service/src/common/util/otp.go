package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateOTP() string {
	max := big.NewInt(1000000)
	otp, err := rand.Int(rand.Reader, max)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%06d", otp.Int64())
}
