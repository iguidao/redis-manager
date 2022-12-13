package useride

import (
	"log"
	"time"

	"github.com/iguidao/redis-manager/src/cfg"

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
	//if err != nil {
	//        log.Println("Error extracting the key", err)
	//}

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		log.Println("Error while signing the token", err)
	}

	//response := Token{tokenString}
	return tokenString
}
