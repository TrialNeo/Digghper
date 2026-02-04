package auth

import (
	"Diggpher/global"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTClaims struct {
	jwt.RegisteredClaims
	UserID uint `json:"userID,omitempty"`
}

func GenerateToken(UserID uint) (string, error) {
	if UserID == 0 {
		return "", errors.New("invalid user ID")
	}

	claims := &JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    global.JwtIssuer,
			Subject:   fmt.Sprintf("%d", UserID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(global.JwtExpiresAt)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: UserID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(global.JwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign JWT: %w", err)
	}

	return tokenString, nil
}

// parseToken 解析并验证 JWT
func parseToken(tokenStr string) (*JWTClaims, error) {
	if tokenStr == "" {
		return nil, errors.New("empty token")
	}
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 确保签名方法是 HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return global.JwtSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}
