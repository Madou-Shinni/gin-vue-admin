package tools

import (
	"github.com/golang-jwt/jwt/v4"
	"strconv"
)

const (
	UserIdKey = "userId"
	ExpKey    = "exp" // 过期时间key
)

// GenToken 生成token map中key=exp过期时间
func GenToken(mapClaims jwt.MapClaims, signed string) (token string, err error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	token, err = claims.SignedString([]byte(signed))
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserIdFromJwt 解析token
func GetUserIdFromJwt(tokenStr string, signed string) (uint, bool) {
	token, err := jwt.ParseWithClaims(tokenStr, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signed), nil
	})
	if err != nil {
		return 0, false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		return 0, false
	} else {
		userIdStr := claims[UserIdKey].(string)
		userId, err := strconv.ParseUint(userIdStr, 10, 64)
		if err != nil {
			return 0, false
		}
		return uint(userId), true
	}
}
