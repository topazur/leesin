package valid

import (
	"fmt"
	"strings"

	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/cast"
	"github.com/topazur/leesin/pkg/util"
	"golang.org/x/text/language"
)

// FieldItem 包装 validator.FieldError 接口
type FieldItem struct {
	// 单个 validator 错误
	fe validator.FieldError

	// Accept-Language
	lang string
	tag  language.Tag

	// locales 翻译接口
	translator locales.Translator
	// 当前语言对应的语言包
	messageCardinalMap map[string]string
	messageItemSlice   MessageItems
}

func NewFieldItem(lang string) (fi FieldItem) {

	fi = FieldItem{
		lang: lang,
	}
	fi.newLangTag()
	fi.newTranslator()

	return
}

// NewLangTag 将传入的 规范性不确定的语言字符串 转化为 language.Tag
func (fi *FieldItem) newLangTag() {
	fi.tag = language.Make(fi.lang)
}

// NewTranslator 根据 language.Tag 的 Base，得到对应的翻译器及其语言包
func (fi *FieldItem) newTranslator() {

	base, _ := fi.tag.Base()

	switch base.String() {
	case "zh":

		fi.translator = zh.New()
		fi.messageCardinalMap = messageCardinalZh
		fi.messageItemSlice = messageItemsZh

	case "en":

		fi.translator = en.New()
		fi.messageCardinalMap = messageCardinalEn
		fi.messageItemSlice = messageItemsEn

	default:

		fi.translator = en.New()
		fi.messageCardinalMap = messageCardinalEn
		fi.messageItemSlice = messageItemsEn

	}
}

/************************************************************/
/************************************************************/

// C 数字格式及其单位(Cardinal复数)
// `ut.Translator.C(key interface{}, num float64, digits uint64, param string) (string, error)`
func (fi *FieldItem) C(key string, param string) string {
	// 计算小数位数，并根据当前语言格式化数字
	var digits uint64
	if idx := strings.Index(param, "."); idx != -1 {
		digits = uint64(len(param[idx+1:]))
	}
	paramNumFmt := fi.translator.FmtNumber(cast.ToFloat64(param), digits)

	p := util.Placeholder(fi.messageCardinalMap[key])
	return p.Replace(p.IteratorCaseCurlybrace, paramNumFmt)
}

// ResetField 重置 validator.FieldError 字段，防止双重for循环时没有匹配到继续使用上一次的结果
func (fi *FieldItem) ResetField() {
	fi.fe = nil
}

func (fi *FieldItem) NoFieldTag() (t string) {
	base, _ := fi.tag.Base()

	switch base.String() {
	case "zh":
		t = "警告: 翻译字段错误: 在语言包中找不到 '%s' Tag"
	case "en":
		t = "warning: error translating FieldError: tag '%s' cannot find in language package"
	default:
		t = "警告: 翻译字段错误: 在语言包中找不到 '%s' Tag"
	}

	return
}

// Error 将 validator.FieldError 翻译成对应语言的可读性文本
func (fi *FieldItem) Error() string {
	tag := fi.fe.Tag() // eg: iscolor
	// tag := fe.ActualTag() // eg: hexcolor|rgb|rgba|hsl|hsla

	oMainFind, ok := fi.messageItemSlice.Find(tag)

	// 没找到 tag 对应的翻译
	if !ok {
		return fmt.Sprintf(fi.NoFieldTag(), tag)
	}

	/// 普通 tag 翻译，tag 与翻译一一对应
	if oMainFind.translation != "" {
		return oMainFind.T(fi.fe.Field(), fi.fe.Param())
	}

	/// 自定义 tag 翻译，tag 与翻译不一一对应
	subMessageItems := oMainFind.customRegisFunc()
	return oMainFind.customTransFunc(fi.fe, subMessageItems, func(key string, param string) string {
		return fi.C(key, param)
	})
}
