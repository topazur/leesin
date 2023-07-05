package valid

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidInController(t *testing.T) {
	type createAccountRequest struct {
		Phone             string    `json:"phone"`
		Email             *string   `json:"email" binding:"omitempty,email"`
		Username          string    `json:"username" binding:"required,alphanum"`
		Nickname          string    `json:"nickname"`
		Password          string    `json:"password" binding:"required,len=6"`
		PasswordUpdatedAt time.Time `json:"password_updated_at" binding:"gt"`
	}

	var test createAccountRequest

	validate.SetTagName("binding")
	err := validate.Struct(test)
	require.NotNil(t, err)

	errStr, errMap := ConvertError(err, "zh")
	require.Emptyf(t, errStr, "当错误成功断言成validator.ValidationErrors时，不回调用err.Error()")
	require.NotEmptyf(t, errMap, "翻译成map结构")

	fmt.Println(errStr, errMap)
}
