package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/duke-git/lancet/v2/random"
	"github.com/spf13/cast"
	"github.com/stretchr/testify/require"
)

func createPayloadParams(duration time.Duration) (string, CreatePayloadParams) {
	return random.RandString(32), CreatePayloadParams{
		Uid:      cast.ToInt64(random.RandInt(100000, 999999)),
		Duration: duration,
	}
}

// TestInvalidKeyPasetoToken 不合规定的密钥长度
func TestInvalidKeyJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(random.RandString(31))
	require.Error(t, err)
	require.EqualErrorf(t, err, err.Error(), "invalid key size: must be at least %d characters", 32)
	require.Nil(t, maker)
}

// TestJWTMaker 测试`正常情况`下token的生成和校验
func TestJWTMaker(t *testing.T) {
	// 时效一分钟
	secreKey, params := createPayloadParams(time.Minute)
	createdAt := time.Now()
	expiredAt := time.Now().Add(params.Duration)

	// 实例化Maker
	maker, err := NewJWTMaker(secreKey)
	require.NoError(t, err)

	// 创建token
	token, payload, err := maker.CreateToken(params)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	// 校验token
	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.NotZero(t, payload.ID)
	require.Equal(t, params.Uid, payload.Uid)
	// 两个时间之间的差异不超过1s
	require.WithinDuration(t, createdAt, payload.CreatedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
}

// TestExpiredJWTToken 测试`过期情况`下token的生成和校验
func TestExpiredJWTToken(t *testing.T) {
	// 时效负一分钟
	secreKey, params := createPayloadParams(-time.Minute)

	// 实例化Maker
	maker, err := NewJWTMaker(secreKey)
	require.NoError(t, err)

	// 创建token
	token, payload, err := maker.CreateToken(params)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	// 校验过期token
	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}

// TestInvalidJWTTokenAlgNone 测试`None算法标头情况`下token的生成和校验
func TestInvalidJWTTokenAlgNone(t *testing.T) {
	// 时效一分钟；手动创建假token
	secreKey, params := createPayloadParams(time.Minute)
	payload, err := NewPayload(params)
	require.NoError(t, err)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	// UnsafeAllowNoneSignatureType特殊常量仅用于测试
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	// 实例化Maker
	maker, err := NewJWTMaker(secreKey)
	require.NoError(t, err)

	// 校验假token
	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
