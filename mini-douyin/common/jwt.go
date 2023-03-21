package common

import (
	"log"
	"mini-douyin/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	ID       int64
	Username string
	jwt.StandardClaims
}

var jwtSecretKey = []byte(config.JWT_KEY)

func GenerateToken(userId int64, username string) string {
	nowTime := time.Now()
	expireTime := nowTime.Add(config.JWT_EXPIRE_TIME).Unix()
	claims := Claims{
		ID:       userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    "mini-douyin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if token, err := tokenClaims.SignedString(jwtSecretKey); err != nil {
		log.Println("generate token fail!")
		return "fail"
	} else {
		log.Println("generate token success!")
		return token
	}
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
