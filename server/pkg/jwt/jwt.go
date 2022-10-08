package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

var mySecret = []byte("memos")

const TokenExpireDuration = time.Hour * 24 * 365

func keyFunc(_ *jwt.Token) (i interface{}, err error) {
	return mySecret, nil
}

func GetToken(userID int64) (token string, err error) {
	c := MyClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "memos",
		},
	}
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, c).SignedString(mySecret)
	return
}

func ParseToken(tokenString string) (claims *MyClaims, err error) {
	var token *jwt.Token
	claims = new(MyClaims)
	token, err = jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return
	}
	if !token.Valid {
		err = errors.New("invalid token")
	}
	return
}
