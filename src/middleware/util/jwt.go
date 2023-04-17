package util

import (
	"github.com/iguidao/redis-manager/src/cfg"
	"github.com/iguidao/redis-manager/src/middleware/mysql"

	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(cfg.Get_Info_String("secretkey"))

var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = "cfun"
)

type Claims struct {
	UserId   int    `json:"userid"`
	UserName string `json:"username"`
	UserType string `json:"usertype"`
	// UserPhone    int64     `json:"UserPhone"`
	// Password string    `json:"Password"`
	// CreaTime time.Time `json:"CreaTime"`
	jwt.StandardClaims
}

// CreateToken 生成一个token
func CreateToken(claims Claims) (string, error) {
	var SigningKey []byte
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SigningKey)
}

func GenerateToken(Username string, Password string) (string, error) {

	nowTime := time.Now()
	expireTime := nowTime.Add(168 * time.Hour)
	muserinfo := mysql.DB.UserInfo(Username)
	username := muserinfo.UserName
	// userid := strconv.Itoa()
	usertype := muserinfo.UserType
	// CreaTime := muserinfo.Base.CreatedAt
	claims := Claims{
		muserinfo.Base.ID,
		username,
		usertype,
		// UserPhone,
		// Password,
		// CreaTime,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "redis-manager",
		},
	}
	token, err := CreateToken(claims)
	// tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
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

func RefreshToken(tokenString string) (string, error) {
	// var username string
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		// userinfo = mysql.DB.UserInfo(claims.UserPhone)
		// username = userinfo.UserName
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(3 * time.Hour).Unix()
		ctoekn, err := CreateToken(*claims)
		return ctoekn, err
	}
	return "", TokenInvalid
}

func GetUserInfo(tokenString string) (UserJWTInfo, error) {
	// var username string
	var userinfo UserJWTInfo
	var err error
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return userinfo, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		userinfo = UserJWTInfo{
			claims.UserId,
			claims.UserName,
			claims.UserType,
		}
		// muserinfo := mysql.DB.UserInfo(claims.UserPhone)
		// userinfo = UserConverge(claims)
		return userinfo, err
	}
	return userinfo, err
}
