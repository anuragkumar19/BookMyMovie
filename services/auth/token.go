package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math"
	"math/big"
)

// TODO: hash otp
func (s *Auth) generateOTP() (rawOTP string, hash string, err error) {
	bigI, err := rand.Int(rand.Reader, big.NewInt(int64(math.Pow10(6)-1)))
	if err != nil {
		return "", "", err
	}
	otp := fmt.Sprintf("%06d", bigI.Int64())
	return otp, otp, nil
}

func (*Auth) generateRandomToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func (s *Auth) matchOTP(rawOTP string, hash string) bool {
	return rawOTP == hash
}
