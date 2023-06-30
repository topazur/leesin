package dao

import (
	"context"
	"testing"
	"time"

	"github.com/duke-git/lancet/v2/random"
	"github.com/stretchr/testify/require"
	"github.com/topazur/leesin/pkg/null"
)

// createRandomAccount 每次测试都是重新生成数据，以免测试干扰导致失败
func createRandomAccount(t *testing.T) Account {
	phone := random.RandNumeral(11)
	username := random.RandNumeralOrLetter(8)

	arg := CreateAccountParams{
		Phone:        null.NewText(phone, true),
		Email:        null.NewText(phone+"@163.com", true),
		Username:     null.NewText(username, true),
		Nickname:     null.NewText(username, true),
		Password:     null.NewText("password", true),
		PasswordSalt: null.NewText("password_salt", true),
		CreatedBy:    null.NewInt8(-1, true),
		UpdatedBy:    null.NewInt8(-1, false),
		UpdatedAt:    null.NewTimestamptz(time.Now(), false),
		DeletedBy:    null.NewInt8(-1, false),
		DeletedAt:    null.NewTimestamptz(time.Now(), false),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Phone, account.Phone)
	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.Username, account.Username)
	require.Equal(t, arg.Nickname, account.Nickname)

	require.Equal(t, arg.Password, account.Password)
	require.Equal(t, arg.PasswordSalt, account.PasswordSalt)
	// 时间类型零值 - '0001-01-01 00:00:00Z'
	require.True(t, account.PasswordUpdatedAt.Time.IsZero())

	require.Equal(t, arg.CreatedBy, account.CreatedBy)
	// 时间间隔不超过1秒
	require.WithinDuration(t, time.Now(), account.CreatedAt.Time, time.Second)

	require.False(t, account.UpdatedBy.Valid)
	require.False(t, account.UpdatedAt.Valid)
	require.False(t, account.DeletedBy.Valid)
	require.False(t, account.DeletedAt.Valid)

	return account
}

// TestQuerier_CreateAccount 创建用户
func TestQuerier_CreateAccount(t *testing.T) {
	createRandomAccount(t)
}

// TestQuerier_UpdateAccount 更新用户
func TestQuerier_UpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:        account.ID,
		UpdatedBy: null.NewInt8(-1, true),
		UpdatedAt: null.NewTimestamptz(time.Now(), true),
		Password:  null.NewText("TestQuerier_UpdateAccount", true),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, arg.UpdatedBy, updatedAccount.UpdatedBy)
	// 时间间隔不超过1秒
	require.WithinDuration(t, time.Now(), updatedAccount.UpdatedAt.Time, time.Second)
	require.Equal(t, arg.Password, updatedAccount.Password)
}
