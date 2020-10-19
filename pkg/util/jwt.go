package util

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go-learning/pkg/setting"
	"log"
	"time"
)

var jwtSecret = []byte(setting.JwtSecret)


type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username,password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 *time.Hour)

	claims := Claims{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "go-learning",
		},
	}
	log.Println("获取username")
	log.Println(claims.Username)
	log.Println(claims.Password)
	log.Println(claims)

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	log.Println(tokenClaims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token,err

}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}