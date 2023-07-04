package util

import (
	"fmt"
	"strings"

	"github.com/spf13/cast"
)

// Placeholder 是string类型别名
type Placeholder string

// String 类型别名转换成string类型
func (p Placeholder) String() string {
	return string(p)
}

// Replace 循环参数，只有当前迭代器的返回值与源串有匹配才会执行替换
func (p Placeholder) Replace(iterator func(index int, arg any) string, args ...any) (result string) {

	// strings.Replace 不会改变原串
	// 先赋值方便循环时只使用 result 变量, 不再需要 condition 变量
	result = p.String()

	for index, arg := range args {
		old := iterator(index, arg)
		result = strings.Replace(result, old, cast.ToString(arg), 1)
	}

	return
}

/************************************************************/
/************************* iterator *************************/
/************************************************************/

func (p Placeholder) IteratorCaseCurlybrace(index int, arg interface{}) string {
	return fmt.Sprintf("{%v}", index)
}

func (p Placeholder) IteratorCaseCurlybraceExcludeZero(index int, arg interface{}) string {
	if index == 0 {
		// 目的是不替换 "{0}" 这个placeholder
		return "{ExcludeZero}"
	}

	return fmt.Sprintf("{%v}", index)
}
