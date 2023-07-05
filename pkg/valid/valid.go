package valid

import (
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/go-playground/validator/v10"
)

func ConvertError(err error, lang string) (string, map[string]string) {
	/// ⏬️ 获取 validator.InvalidValidationError 类型的errors
	/// eg: "validator: (nil " + e.Type.String() + ")"
	invalidErr, ok := err.(*validator.InvalidValidationError)
	if ok {
		return invalidErr.Error(), nil
	}

	/// ⏬️ 获取 validator.ValidationErrors 类型的errors
	/// eg: "Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag" => 转化为 map[string]FieldError
	validErrs, ok := err.(validator.ValidationErrors)

	// 存储为map格式
	if ok {
		errMap := make(map[string]string)
		fi := NewFieldItem(lang)

		for _, fe := range validErrs {
			// 结构体大驼峰转化为蛇形命名法, 便于json返回与前端交互
			key := strutil.SnakeCase(fe.Field())

			// 每次循环对 validator.FieldError 重新赋值，并得到其翻译结果
			fi.fe = fe
			errMap[key] = fi.Error()
		}
		return "", errMap
	}

	/// ⏬️ 其他错误
	return err.Error(), nil
}
