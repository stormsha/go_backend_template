package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"stormsha.com/gbt/model"
	"time"
)

type User struct {
	Id          int    `json:"id"`
	UserAccount string `json:"user_account"`
	UserName    string `json:"user_name"`
}

type TokenClaims struct {
	Id          int    `json:"id"`
	UserAccount string `json:"user_account"`
	UserName    string `json:"user_name"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(conf.JwtSecretKey) // JWT 密钥

func CheckUserPermission(tokenString string) (*User, error) {
	claims := &TokenClaims{}
	token, err := jwt.ParseWithClaims(tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		},
	)
	if err != nil || !token.Valid {
		if err == nil {
			err = errors.New("token 解析失败")
		}
		return nil, err
	}
	user := &User{
		Id:          claims.Id,
		UserAccount: claims.UserAccount,
		UserName:    claims.UserName,
	}
	return user, nil
}

func GetUserToken(user *model.User) (string, error) {
	// 生成 JWT token
	claims := TokenClaims{
		Id:          user.ID,
		UserAccount: user.UserAccount,
		UserName:    user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token 有效期为 24 小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	return tokenString, err
}
