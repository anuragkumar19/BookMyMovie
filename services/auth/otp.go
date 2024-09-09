package auth

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

// TODO: hash otp
func (*Auth) generateOTP() (rawOTP string, hash string, err error) {
	bigI, err := rand.Int(rand.Reader, big.NewInt(int64(math.Pow10(6)-1)))
	if err != nil {
		return "", "", err
	}
	otp := fmt.Sprintf("%06d", bigI.Int64())
	return otp, otp, nil
}

func (*Auth) matchOTP(rawOTP string, hash string) bool {
	return rawOTP == hash
}
