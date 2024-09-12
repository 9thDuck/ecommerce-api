package auth

import (
	"errors"
	"time"

	"github.com/9thDuck/ecommerce-api.git/common"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type TokenClaims struct {
	ID   uuid.UUID `json:"id"`
	Role int       `json:"role"`
}

func GenerateToken(claims TokenClaims) (accessTokenStr string, refreshTokenStr string, err error) {
	EXPIRY_ACCESS_TOKEN_VALUE_IN_MINUTES_IN_UNIX := time.Now().Add(common.APP_CONFIG.GetExpiryAccessTokenDurationInMinutes()).Unix()
	EXPIRY_REFRESH_TOKEN_DURATION_IN_HOURS_IN_UNIX := time.Now().Add(common.APP_CONFIG.GetExpiryRefreshTokenDurationInHours()).Unix()

	accessTokenClaims := jwt.MapClaims{
		"id":   claims.ID,
		"role": claims.Role,
		"exp":  EXPIRY_ACCESS_TOKEN_VALUE_IN_MINUTES_IN_UNIX,
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenStr, err = accessToken.SignedString([]byte(common.APP_CONFIG.GetJwtSecret()))

	if err != nil {
		return "", "", err
	}

	refreshTokenClaims := jwt.MapClaims{
		"id":   claims.ID,
		"role": claims.Role,
		"exp":  EXPIRY_REFRESH_TOKEN_DURATION_IN_HOURS_IN_UNIX,
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)

	refreshTokenStr, err = refreshToken.SignedString([]byte(common.APP_CONFIG.GetJwtSecret()))
	if err != nil {
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected token signing method detected")
		}
		return []byte(common.APP_CONFIG.GetJwtSecret()), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}