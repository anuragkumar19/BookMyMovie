package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"time"

	"bookmymovie.app/bookmymovie/database"
	"bookmymovie.app/bookmymovie/services/serviceserrors"
	"github.com/golang-jwt/jwt/v5"
)

var (
	errTokenInvalid = errors.New("invalid token")
)

type Metadata struct {
	userID         int64
	userRole       database.Roles
	refreshTokenID int64

	valid bool
	err   error
}

func (m *Metadata) Valid() error {
	if !m.valid {
		if nil == m.err {
			return errors.New("auth.Metadata must be created from functions available in auth package")
		}
		return serviceserrors.New(serviceserrors.ErrorConflict, m.err.Error())
	}
	return nil
}

func (m *Metadata) UserID() int64 {
	return m.userID
}

func (m *Metadata) RefreshTokenID() int64 {
	return m.refreshTokenID
}

func (m *Metadata) UserRole() database.Roles {
	return m.userRole
}

func (s *Auth) GetMetadata(accessToken string) Metadata {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(s.config.AccessTokenSecret), nil
	})
	if err != nil {
		return Metadata{
			valid: false,
			err:   err,
		}
	}
	if !token.Valid {
		return Metadata{
			valid: false,
			err:   errTokenInvalid,
		}
	}
	claims := token.Claims.(jwt.MapClaims) //nolint:errorlint,errcheck

	id := int64(claims["id"].(float64))                      //nolint:errorlint,errcheck
	userID := int64(claims["user_id"].(float64))             //nolint:errorlint,errcheck
	userRole := database.Roles(claims["user_role"].(string)) //nolint:errorlint,errcheck

	if _, ok := s.revokedTokens[id]; ok {
		return Metadata{
			valid: false,
			err:   errTokenInvalid,
		}
	}

	return Metadata{
		refreshTokenID: id,
		userID:         userID,
		userRole:       userRole,
		valid:          true,
	}
}

func (s *Auth) startBackgroundRevokedTokenCleanup() {
	go func() {
		for {
			for k, v := range s.revokedTokens {
				if v.Before(time.Now()) {
					delete(s.revokedTokens, k)
				}
			}
			time.Sleep(s.config.AccessTokenLifetime)
		}
	}()
}

func (*Auth) generateRandomToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b) + strconv.Itoa(int(time.Now().UnixNano())), nil
}

func (s *Auth) generateAccessToken(rt *database.RefreshToken) (token string, expiry time.Time, err error) {
	t := jwt.New(jwt.SigningMethodHS256)

	now := time.Now()
	expiry = now.Add(s.config.AccessTokenLifetime)

	claims := t.Claims.(jwt.MapClaims) //nolint:errorlint,errcheck

	claims["iss"] = s.config.AppPublicHost
	claims["id"] = rt.ID
	claims["user_id"] = rt.UserID
	claims["user_role"] = rt.UserRole
	claims["iat"] = now.Unix()
	claims["exp"] = expiry.Unix()

	st, err := t.SignedString([]byte(s.config.AccessTokenSecret))
	return st, expiry, err
}
