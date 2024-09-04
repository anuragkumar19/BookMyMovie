package auth

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
)

var (
	ErrTokenInvalid = errors.New("invalid token")
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
