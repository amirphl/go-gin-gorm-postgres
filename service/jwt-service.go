package service

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

// JWTService ...
type JWTService interface {
	GenerateToken(userID uint64) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID uint64 `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func (j *jwtService) GenerateToken(userID uint64) string {
	claims := &jwtCustomClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))

	if err != nil {
		panic(err)
	}

	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	f := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}

		return []byte(j.secretKey), nil
	}

	return jwt.Parse(token, f)
}

// CreateJWTService ...
func CreateJWTService() JWTService {
	secKey := os.Getenv("JWT_SECRET_KEY")

	if secKey == "" {
		panic("JWT secret key is nil")
	}

	return &jwtService{
		secretKey: secKey,
		issuer:    os.Getenv("JWT_ISSUER"),
	}
}
