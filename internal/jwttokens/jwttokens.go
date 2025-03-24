package jwttokens

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	timePeriod = 30 * 24 * time.Hour
)

var (
	secretKey = []byte("k2Sf3Y1@z\rG#1hX7%%d4Q,9$l")
)

type userIdClaims struct {
	jwt.MapClaims
	UserId int64
}

func NewUserIdToken(userId int64) (string, error) {
	claims := userIdClaims{
		MapClaims: jwt.MapClaims{
			"exp": time.Now().Add(timePeriod).Unix(),
		},
		UserId: userId,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := accessToken.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}

func ParseUserIdToken(token string) (int64, error) {
	accessToken, err := jwt.ParseWithClaims(token, &userIdClaims{}, keyFunc)
	if err != nil {
		return -1, err
	}

	if !accessToken.Valid {
		return -1, errors.New("token is invalid")
	}

	claims, ok := accessToken.Claims.(*userIdClaims)
	if !ok {
		return -1, errors.New("token claims are not of type *userIdClaims")
	}

	return claims.UserId, nil
}

func keyFunc(accessToken *jwt.Token) (interface{}, error) {
	_, ok := accessToken.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("invalid signing method")
	}

	return secretKey, nil
}
