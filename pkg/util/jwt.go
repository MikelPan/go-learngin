package util

import (
	//"github.com/MikelPan/go-learning/pkg/setting"
	jwt "github.com/dgrijalva/jwt-go"
)

//var jwtSecret = []byte(setting.JwtSecret)


type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(username,password string) (string, error) {
	//nowTime := time.Now()
	//expireTime := nowTime.Add(3 *time.Hour)

	//claims := Claims{
	//	username,
	//	password,
	//	jwt.StandardClaims {
	//		NotBefore: int64(time.Now().Unix() - 1000),
	//		ExpiresAt: expireTime.Unix(),
	//		Issuer: "go-learning",
	//	},
	//}
	//tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//token, err := tokenClaims.SignedString(jwtSecret)
	/**if err != nil {
		log.Println(err)
	}
	*/
	return "", nil

}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}