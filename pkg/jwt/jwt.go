package jwt

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWT struct{}

func (JWT) CreateAccessToken(id, name, username string) (string, error) {
	payload := &AccessPayload{ID: uuid.MustParse(id), Name: name, Username: username}
	signingKey := []byte(os.Getenv("SIGNKEY"))
	expireTime := time.Now().Add(1 * time.Hour)

	payload.ExpiresAt = jwt.NewNumericDate(expireTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(signingKey)
	if err != nil {
		log.Printf("ERROR CreateAccessToken fatal error: %v", err)
		return "", err
	}

	return tokenStr, nil
}

func (JWT) CreateRefreshToken(id string) (string, error) {
	payload := &RefreshPayload{ID: uuid.MustParse(id), IsRT: true}
	signingKey := []byte(os.Getenv("SIGNKEY"))
	expireTime := time.Now().AddDate(0, 0, 30)

	payload.ExpiresAt = jwt.NewNumericDate(expireTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(signingKey)
	if err != nil {
		log.Printf("ERROR CreateRefreshTokenFatalError: %v", err)
		return "", err
	}

	return tokenStr, nil
}

func (JWT) ValidateAccessToken(token string) (string, error) {
	signingKey := []byte(os.Getenv("SIGNKEY"))
	claims := AccessPayload{}

	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (any, error) {
		return signingKey, nil
	})
	if err != nil {
		log.Printf("ERROR Validate fatal error: %v", err)
		return "", nil
	}

	if !parsedToken.Valid {
		log.Println("ERROR VallidateAccessToken fatal error: invalid access token")
		return "", errors.New("invalid refresh token")
	}

	return claims.ID.String(), nil
}

func (JWT) ValidateRefreshToken(token string) (string, error) {
	signingKey := []byte(os.Getenv("SIGNKEY"))
	claims := RefreshPayload{}

	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (any, error) {
		return signingKey, nil
	})
	if err != nil {
		log.Printf("ERROR Validate fatal error: %v", err)
		return "", nil
	}

	if !parsedToken.Valid {
		log.Println("ERROR VallidateRefreshToken fatal error: invalid refresh token")
		return "", errors.New("invalid refresh token")
	}

	if !claims.IsRT {
		log.Println("ERROR ValidateRefreshToken fatal error: invalid refresh token")
		return "", errors.New("invalid refresh token")
	}

	return claims.ID.String(), nil
}
