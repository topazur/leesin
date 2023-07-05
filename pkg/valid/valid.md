## `package valid` 的目的是翻译校验错误

- [李文周 - 翻译校验错误 & 自定义字段校验](https://www.liwenzhou.com/posts/Go/validator-usages/)

- [validator提供的翻译插件](https://github.com/go-playground/validator/v10/translations)

- 翻译插件选型

  ```go
  // "golang.org/x/text/language"
  // 解析字符串得到标准language.Tag，更加安全准确
  english := language.Make("en") <=> language.English.String()

  // 方式一:  "golang.org/x/text/message"
  // 缺点: Sprintf中的参数format模板不能拼接, 不能确定小数位数
  // "%d" -> 整数；"%.2f" -> 保留两位小数
  p := message.NewPrinter(language.Chinese)
  num := 1000000
  formatted := p.Sprintf("%d", num)

  // "github.com/gin-gonic/gin"
  // "github.com/gin-gonic/gin/binding"
  // "github.com/go-playground/locales/en"
  // "github.com/go-playground/locales/zh"
  // ut "github.com/go-playground/universal-translator"
  // "github.com/go-playground/validator/v10"
  // enTranslations "github.com/go-playground/validator/v10/translations/en"
  // zhTranslations "github.com/go-playground/validator/v10/translations/zh"
  // 定义全局翻译器代价过大
  locales.translator.FmtNumber


  => 最终自己模仿 `github.com/go-playground/validator/v10/translations` 实现翻译校验错误
  ```

<br />

### validator.FieldError 接口转结构体 (已放弃此方案，直接将FieldError整体存入FieldItem中)

```go
type FieldError struct {
	Tag       string `json:"tag"`
	ActualTag string `json:"actualTag"`

	/// Namespace
	/// StructNamespace
	/// Field
	/// StructField

	Value interface{} `json:"value"`
	Param string      `json:"param"`

	/**
	 * 区别
	 * Kind returns the Field's reflect Kind => eg. time.Time's kind is a struct
	 * Type returns the Field's reflect Type => eg. time.Time's type is time.Time
	 */
	/// Kind
	Type string `json:"type"`

	/// Error
}
```

<br />

## 坑

```go
- v9 与 v10 版本
  github.com/go-playground/validator v9 
  err.(validator.ValidationErrors) 断言得到 validator.ValidationErrors(nil)，但实际是有错误存在

- Bind
  ctx.ShouldBind 方法用于将请求数据根据 Content-Type 自动进行解析并绑定到结构体。它可以根据请求头中的 Content-Type 自动选择解析方法，支持表单数据、JSON、XML 等格式。
  ctx.ShouldBindJSON 方法是 ctx.ShouldBind 的特定版本，用于将 JSON 格式的请求数据绑定到结构体。它会自动解析请求主体中的 JSON 数据并进行绑定。
```

