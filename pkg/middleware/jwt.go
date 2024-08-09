package middleware

import (
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID string `json:"userId"`
	jwt.RegisteredClaims
}

func SignAccessToken(userId string) (string, error) {
	claims := CustomClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Second)),
			Issuer:    "pickurpage.com",
			Audience:  jwt.ClaimStrings{userId},
		},
	}

	secret := os.Getenv("ACCESS_TOKEN_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyAccessToken(c *gin.Context) (*CustomClaims, error) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return nil, errors.New("Unauthorized")
	}

	bearerToken := authHeader[len("Bearer "):]

	token, err := jwt.ParseWithClaims(bearerToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("Unauthorized")
		}
		return nil, errors.New("Bad Request")
	}

	if !token.Valid {
		return nil, errors.New("Unauthorized")
	}

	claims, ok := token.Claims.(*CustomClaims)

	if !ok {
		return nil, errors.New("Unauthorized")
	}
	return claims, nil
}

func SignRefreshToken(userId string) (string, error) {
	claims := CustomClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(365 * 24 * time.Hour)),
			Issuer:    "pickurpage.com",
			Audience:  jwt.ClaimStrings{userId},
		},
	}

	secret := os.Getenv("REFRESH_TOKEN_SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyRefreshToken(refreshToken string) (string, error) {
	token, err := jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("REFRESH_TOKEN_SECRET")), nil
	})

	if err != nil {
		return "", errors.New("unauthorized")
	}

	if !token.Valid {
		return "", errors.New("unauthorized")
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return "", errors.New("unauthorized")
	}

	userID := claims.Audience[0]
	if err != nil {
		return "", errors.New("unauthorized")
	}

	return userID, nil
}
