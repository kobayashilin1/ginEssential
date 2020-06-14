package common

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kobayashilin1/ginEssential/model"
	"time"
)

var jwtkey = []byte("a_secrect_crect")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}


func ReleaseToken(user model.User)(string, error){
	expirationTime := time.Now().Add(7 * 24 * time.Hour)//设置token的有限时间：7d * 24h * time
	claims := &Claims{
		UserId:			user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:expirationTime.Unix(),
			IssuedAt: time.Now().Unix(),//toke发放的时间
			Issuer:"kobayashilin1",//token发放者
			Subject: "user token",
		},
		//token由三部分组成：header.payload.hash生成的由前面两部分(header+payload)+key三者共同生成的值。

	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err := token.SignedString(jwtkey)

	if  err != nil {
		return "", err
	}
	return tokenString, nil
}

//从tokenString中解析出claims并返回。
func ParseToken(tokenString string)(*jwt.Token, *Claims, error){
	claims := &Claims{}
	
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})

	return token, claims, err
}