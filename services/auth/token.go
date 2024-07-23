package auth

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"math"
	"math/big"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"github.com/golang-jwt/jwt/v5"
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

func (s *Auth) generateRefreshToken(device *database.RefreshToken) (token string, err error) {
	t := jwt.New(jwt.SigningMethodHS256)

	claims := t.Claims.(jwt.MapClaims)

	claims["iss"] = s.config.Host
	claims["id"] = device.ID
	claims["user_id"] = device.UserID
	claims["user_role"] = device.UserRole
	claims["iat"] = device.CreatedAt.Time.Unix()
	claims["exp"] = device.ExpireAt.Time.Unix()

	return t.SignedString([]byte(s.config.RefreshTokenSecret))
}

func (s *Auth) generateAccessToken(device *database.RefreshToken) (token string, expiry time.Time, err error) {
	t := jwt.New(jwt.SigningMethodHS256)

	now := time.Now()
	expiry = now.Add(s.config.AccessTokenLifetime)

	claims := t.Claims.(jwt.MapClaims)

	claims["iss"] = s.config.Host
	claims["id"] = device.ID
	claims["user_id"] = device.UserID
	claims["user_role"] = device.UserRole
	claims["iat"] = now.Unix()
	claims["exp"] = expiry.Unix()

	st, err := t.SignedString([]byte(s.config.AccessTokenSecret))
	return st, expiry, err
}
