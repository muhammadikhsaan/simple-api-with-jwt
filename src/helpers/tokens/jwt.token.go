package tokens

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//JWTService is interface for JWT
type JWTService interface {
	GenerateToken(JWTData SignetureToken) (string, error)
	ValidateToken(tokenString string) (*jwtCustomClaims, error)
}

type SignetureToken struct {
	Email string `json:"email"`
}

type jwtCustomClaims struct {
	SignetureToken
	jwt.StandardClaims
}

type jwtService struct {
	signer   []byte
	issuer   string
	audience string
}

//NewJWTService declarate JWT value
func NewJWTService() JWTService {
	return &jwtService{
		signer:   []byte(os.Getenv("JWT_PASSPHRASE")),
		issuer:   os.Getenv("JWT_ISSUER"),
		audience: os.Getenv("JWT_AUDIENCE"),
	}
}

func (j *jwtService) GenerateToken(JWTData SignetureToken) (string, error) {
	claims := &jwtCustomClaims{
		JWTData,
		jwt.StandardClaims{
			Audience: j.audience,
			Issuer:   j.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString(j.signer)

	if err != nil {
		return "", err
	}

	return ss, nil
}

func (j *jwtService) ValidateToken(tokenString string) (*jwtCustomClaims, error) {

	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	token, err := jwt.ParseWithClaims(tokenString, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.signer, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*jwtCustomClaims)

	if !ok && !token.Valid {
		return nil, errors.New("token invalid")
	}

	return claims, nil
}
