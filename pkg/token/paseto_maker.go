package token

import (
	"fmt"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

// PasetoMaker is Maker interface implement
type PasetoMaker struct {
	// 使用 v2 版本
	paseto *paseto.V2
	// 对称密钥 (无第三方校验)
	symmetricKey []byte
}

// CreateToken is implement token.Maker
func (p *PasetoMaker) CreateToken(params CreatePayloadParams) (string, *Payload, error) {
	payload, err := NewPayload(params)
	if err != nil {
		return "", payload, err
	}

	// 传入加密算法和有效载荷和可选页脚 生成token
	token, err := p.paseto.Encrypt(p.symmetricKey, payload, nil)
	return token, payload, err
}

// VerifyToken is implement token.Maker
func (p *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	// 空结构体接受解码后的有效载体
	payload := &Payload{}

	// 解密操作
	err := p.paseto.Decrypt(token, p.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// payload实现的Valid方法，检验是否过期
	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// NewPasetoMaker creates a nwe PasetoMaker
func NewPasetoMaker(symmetricKey string) (Maker, error) {
	// 判断 chacha poly cipher 算法需要的密钥长度是否相等
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}
