package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/emelyanenkoig/todo"
	"github.com/emelyanenkoig/todo/pkg/repository"
)

const (
	salt       = "sdfkaskfkrl"
	signingKey = "asdfghj"
	tokenTTL   = 6 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = GeneratePasswordHash(user.Password)
	return as.repo.CreateUser(user)

}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (as *AuthService) GenerateToken(username string, password string) (string, error) {
	user, err := as.repo.GetUser(username, GeneratePasswordHash(password))

	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
	},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))

}

func (as *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims not a type *tokenClaims")
	}
	return claims.UserID, nil
}
