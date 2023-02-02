package useride

import (
	"time"

	"github.com/iguidao/redis-manager/src/cfg"
	"github.com/iguidao/redis-manager/src/middleware/logger"

	jwt_lib "github.com/golang-jwt/jwt/v4"
)

type Token struct {
	Token string `json:"token"`
}

func Token_Get() string {
	SecretKey := cfg.Get_Info_String("secretkey")
	token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
	claims := make(jwt_lib.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		logger.Error("Error while signing the token", err)
	}

	//response := Token{tokenString}
	return tokenString
}
