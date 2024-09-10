package auth

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"

	"golang.org/x/crypto/bcrypt"
)

func (*Auth) generateOTP() (rawOTP string, hash string, err error) {
	bigI, err := rand.Int(rand.Reader, big.NewInt(int64(math.Pow10(6)-1)))
	if err != nil {
		return "", "", err
	}
	otp := fmt.Sprintf("%06d", bigI.Int64())

	b, err := bcrypt.GenerateFromPassword([]byte(otp), 12)
	if err != nil {
		return "", "", err
	}
	return otp, string(b), nil
}

func (*Auth) matchOTP(rawOTP string, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(rawOTP)) == nil
}
