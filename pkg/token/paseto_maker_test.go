package token

import (
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/random"
	"github.com/stretchr/testify/require"
)

// TestInvalidKeyPasetoToken 不合规定的密钥长度
func TestInvalidKeyPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(random.RandString(31))
	require.Error(t, err)
	require.EqualError(t, err, "invalid key size: must be exactly 32 characters")
	require.Nil(t, maker)
}

// TestPasetoMaker 测试`正常情况`下token的生成和校验
func TestPasetoMaker(t *testing.T) {
	// 时效一分钟
	secreKey, params := createPayloadParams(time.Minute)
	createdAt := time.Now()
	expiredAt := time.Now().Add(params.Duration)

	// 实例化Maker
	maker, err := NewPasetoMaker(secreKey)
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

// TestExpiredPasetoToken 测试`过期情况`下token的生成和校验
func TestExpiredPasetoToken(t *testing.T) {
	// 时效负一分钟
	secreKey, params := createPayloadParams(-time.Minute)

	// 实例化Maker
	maker, err := NewPasetoMaker(secreKey)
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

// TestInvalidPasetoToken paseto中不存在None算法 (修改token字符串模拟Invalid情况)
func TestInvalidPasetoToken(t *testing.T) {
	// 时效一分钟
	secreKey, params := createPayloadParams(time.Minute)

	// 实例化Maker
	maker, err := NewPasetoMaker(secreKey)
	require.NoError(t, err)

	// 创建token
	token, payload, err := maker.CreateToken(params)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	// 检验假token
	payload, err = maker.VerifyToken(token + "fake")
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
