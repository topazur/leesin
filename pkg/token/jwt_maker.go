package token

import (
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

// JWTMaker is Maker interface implement
type JWTMaker struct {
	secretKey string
}

// CreateToken is implement token.Maker
func (j *JWTMaker) CreateToken(info CreatePayloadParams) (string, *Payload, error) {
	payload, err := NewPayload(info)
	if err != nil {
		return "", payload, err
	}

	// 传入加密算法和有效载荷 (并且通过payload中的Valid方法判断过期时间，若晚于当前时间没有加密的必要)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	// 根据 secretKey 生成 token
	token, err := jwtToken.SignedString([]byte(j.secretKey))
	return token, payload, err
}

// VerifyToken is implement token.Maker
func (j *JWTMaker) VerifyToken(token string) (*Payload, error) {
	// 回调函数
	var keyFunc jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) {
		// SigningMethodHS256是SigningMethodHMAC的一个实例
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			//	token加密算法与我们不同，可能是伪造的
			return nil, ErrInvalidToken
		}
		return []byte(j.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	// 根据 解析返回的错误的Inner 做一下错误类型的判断
	if err != nil {
		vErr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(vErr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	// 解析成功：把jwtToken.Claims转换为Payload对象
	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}
	return payload, nil
}

// NewJWTMaker creates a nwe JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}

	return &JWTMaker{secretKey: secretKey}, nil
}
