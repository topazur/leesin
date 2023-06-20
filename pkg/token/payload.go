package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by th VerifyToken function
var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token has invalid")

	// 请求头FieldName: field会自动转化为小写，因此不用担心处理大小写兼容问题
	AuthorizationHeaderKey = "authorization"
	// totken前缀
	AuthorizationTypeBearer = "bearer"
	// 中间件内部传递参数的key
	AuthorizationPayloadKey = "authorization_payload"
)

type CreatePayloadParams struct {
	Uid      int64         `json:"uid"`      // 用户信息-uid
	Duration time.Duration `json:"duration"` // 有效时间间隔
}

// Payload contains th payload data of the token
type Payload struct {
	// 匿名结构体在声明时需要带上结构体名，但是使用时可以直接访问嵌套结构体内的字段
	CreatePayloadParams

	ID        uuid.UUID `json:"id"`         // token专属id (机制：唯一标识，仅限含uuid的token生效)
	CreatedAt time.Time `json:"created_at"` // 创建时间
	ExpiredAt time.Time `json:"expired_at"` // 过期时间
}

// Valid is implement jwt.Claims
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

// NewPayload creates a new token payload with username and duration
func NewPayload(params CreatePayloadParams) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		CreatePayloadParams: params,
		ID:                  id,
		CreatedAt:           time.Now(),
		ExpiredAt:           time.Now().Add(params.Duration),
	}
	return payload, nil
}
