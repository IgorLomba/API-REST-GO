package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// issure é quem assinou o token
type jwtService struct {
	secretKey string
	issure    string
}

func NewJwtService() *jwtService {
	return &jwtService{
		// o certo seria os.Getenv(Sha256_encoder("secret-key"))
		secretKey: "secret-key",
		issure:    "person-api",
	}
}

type Claim struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

func (s *jwtService) GenerateToken(id uint) (string, error) {
	claim := &Claim{
		id,
		jwt.StandardClaims{
			//  token will expire in 2 hours
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issure,
			// created at...
			IssuedAt: time.Now().Unix(),
		},
	}
	// calling token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// signing token
	t, err := token.SignedString([]byte(s.secretKey))

	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %v", token)
		}

		return []byte(s.secretKey), nil
	})

	return err == nil
}
