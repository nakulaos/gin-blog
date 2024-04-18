package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"gvb_server/global"
	"time"
)

// JWT的payload 中的数据
type JwtPayLoad struct {
	//Username string `json:"username"` //
	Nickname string `json:"nickname"` //
	Role     int    `json:"role"`     //
	UserID   uint   `json:"user_id"`  //
}
type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

// 生成Token
func GenToken(jwtPayLoad JwtPayLoad) (string, error) {
	MySecret := []byte(global.Config.Jwt.Secret)
	claim := CustomClaims{
		jwtPayLoad,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(global.Config.Jwt.Expires)).Unix(), //2个小时过期
			Issuer:    global.Config.Jwt.Issuer,                                                    //签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

// 解析token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	MySecret := []byte(global.Config.Jwt.Secret)

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Log.Error(err.Error())
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("token invalid！")
}
