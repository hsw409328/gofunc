/**
 * Author: haoshuaiwei 
 * Date: 2019-02-27 17:19 
 */

package go_jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	SignKey          = "xxxxx"
	TokenExpired     = errors.New("令牌已经过期")
	TokenNotValidYet = errors.New("不是有效的令牌")
	TokenMalformed   = errors.New("令牌发生异常")
	TokenInvalid     = errors.New("令牌无效")
)

func JwtSetSignKey(key string) {
	SignKey = key
}

// 载荷，可以加一些自己需要的信息
type CustomClaimsPayload struct {
	Host string
	jwt.StandardClaims
}

type JwtObject struct {
	SignKey []byte
}

func NewJwtObject() *JwtObject {
	return &JwtObject{
		SignKey: []byte(SignKey),
	}
}

func (c *JwtObject) CreateJwtToken(claims CustomClaimsPayload) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(c.SignKey)
}

func (c *JwtObject) JwtParse(tokenString string) (*CustomClaimsPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsPayload{}, func(token *jwt.Token) (interface{}, error) {
		return c.SignKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&(jwt.ValidationErrorExpired) != 0 {
				return nil, TokenExpired
			} else if ve.Errors&(jwt.ValidationErrorNotValidYet) != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaimsPayload); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

func (c *JwtObject) JwtRefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	// 解析
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsPayload{}, func(token *jwt.Token) (interface{}, error) {
		return c.SignKey, nil
	})
	if err != nil {
		return "", err
	}
	// 判断payload
	if claims, ok := token.Claims.(*CustomClaimsPayload); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		// 延长时间，重新创建
		return c.CreateJwtToken(*claims)
	}
	return "", TokenInvalid
}
