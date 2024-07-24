package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrTokenInvalid = errors.New("invalid token")
)

// type tokenData struct {
// 	iss      string
// 	id       int64
// 	userId   int64
// 	userRole database.Roles
// }

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
	return base64.RawURLEncoding.EncodeToString(b) + strconv.Itoa(int(time.Now().UnixNano())), nil
}

func (s *Auth) matchOTP(rawOTP string, hash string) bool {
	return rawOTP == hash
}

func (s *Auth) generateAccessToken(rt *database.RefreshToken) (token string, expiry time.Time, err error) {
	t := jwt.New(jwt.SigningMethodHS256)

	now := time.Now()
	expiry = now.Add(s.config.AccessTokenLifetime)

	claims := t.Claims.(jwt.MapClaims)

	claims["iss"] = s.config.Host
	claims["id"] = rt.ID
	claims["user_id"] = rt.UserID
	claims["user_role"] = rt.UserRole
	claims["iat"] = now.Unix()
	claims["exp"] = expiry.Unix()

	st, err := t.SignedString([]byte(s.config.AccessTokenSecret))
	return st, expiry, err
}

// func parseAndVerifyToken(tokenStr string, secret string) (tokenData, error) {
// 	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
// 		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
// 		}
// 		return []byte(secret), nil
// 	})
// 	if err != nil {
// 		return tokenData{}, err
// 	}
// 	if !token.Valid {
// 		return tokenData{}, ErrTokenInvalid
// 	}
// 	claims := token.Claims.(jwt.MapClaims)
// 	iss, err := claims.GetIssuer()
// 	if err != nil {
// 		return tokenData{}, err
// 	}
// 	id := claims["id"].(int64)
// 	userId := claims["user_id"].(int64)
// 	userRole := claims["user_role"].(database.Roles)
// 	return tokenData{iss: iss, id: id, userId: userId, userRole: userRole}, nil
// }
