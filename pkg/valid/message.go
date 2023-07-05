package valid

import (
	"os"

	"github.com/duke-git/lancet/v2/slice"
	"github.com/go-playground/validator/v10"
	"github.com/topazur/leesin/pkg/util"
)

// FmtCardinal 为了简化函数类型书写难度
type FmtCardinal func(key string, param string) string

// 每个 validator Tag 对应的语义化翻译值
type MessageItem struct {
	tag             string
	translation     string
	override        bool
	customRegisFunc func() MessageItems
	customTransFunc func(fe validator.FieldError, items MessageItems, fmtCardinal FmtCardinal) string
}

// Deprecated: This function is deprecated. Use NewFunction FieldItem.C instead.
// C 数字格式及其单位(Cardinal复数)
// `ut.Translator.C(key interface{}, num float64, digits uint64, param string) (string, error)`
func (item MessageItem) C(key string, param string) string {
	return ""
}

// T 通过key找到item之后，调用T简单替换参数
// `ut.Translator.T(key interface{}, params ...string) (string, error)`
// @DESC 虽然明确知道参数params的类型是string，但是为了兼容StringReplacePlaceholder函数，还是使用any类型。否则因类型不同调用时不会被扩展运算符展开
func (item MessageItem) T(params ...any) string {

	p := util.Placeholder(item.translation)

	// `go:build test`
	if os.Getenv("VALIDATOR_REPLACE_FIELD") == "onlytest" {
		return p.Replace(p.IteratorCaseCurlybrace, params...)
	}

	// `go:build !test` 返回前端时保留 `{0}` 占位符，方便前端替换成字段名或者其他label
	return p.Replace(p.IteratorCaseCurlybraceExcludeZero, params...)
}

/************************************************************/
/************************************************************/

// MessageItems 是 []MessageItem 切片的类型别名 (invalid receiver type []MessageItem)
type MessageItems []MessageItem

// Add 快速创建MessageItem并append到切片中
// NOTICE: 指针接收器才可以append
func (items *MessageItems) Add(tag string, translation string, override bool) {
	*items = append(*items, MessageItem{tag: tag, translation: translation, override: override})
}

// 迭代slice的元素, 查找当前切片内存在的 validator Tag.
func (items MessageItems) Find(validatorTag string) (MessageItem, bool) {
	return slice.FindBy(items, func(index int, item MessageItem) bool {
		return item.tag == validatorTag
	})
}
