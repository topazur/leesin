package valid

import (
	"os"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

var validate *validator.Validate

// TestMain 同一目录下的所有测试之前执行，可做一些初始化操作
func TestMain(m *testing.M) {
	validate = validator.New()

	// `go-playground/validator`包默认使用标记名称“validate”，可通过SetTagName来更改标记名称
	// https://github.com/gin-gonic/gin/blob/51aea73ba0f125f6cacc3b4b695efdf21d9c634f/binding/default_validator.go#L95
	validate.SetTagName("validate")

	// NOTICE: 标识仅测试环境替换字段名，否则会保存 `{0}` 占位符
	os.Setenv("VALIDATOR_REPLACE_FIELD", "onlytest")

	// 🌈 运行单元测试函数
	exitCode := m.Run()

	// 退出测试 (退出之前可以完成一些清理操作)
	validate = nil
	os.Exit(exitCode)
}

type testCase struct {
	ns       string
	expected string
}

func loopTestFunc(
	t *testing.T,
	errs validator.ValidationErrors,
	tests []testCase,
	lang string,
) {

	fi := NewFieldItem(lang)

	for _, tt := range tests {
		fi.ResetField()

		for _, e := range errs {
			if tt.ns == e.Namespace() {
				fi.fe = e
				break
			}
		}

		require.NotNilf(t, fi.fe, "没找到字段对应的validator.FieldError？")
		require.Equal(t, tt.expected, fi.Error())
	}
}
