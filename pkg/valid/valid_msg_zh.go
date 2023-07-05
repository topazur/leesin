package valid

import (
	"fmt"
	"reflect"
	"time"

	"github.com/go-playground/validator/v10"
)

// 检验器的复数翻译
var messageCardinalZh = map[string]string{
	"character":   "{0}个字符",
	"items-item":  "{0}项",
	"placeholder": "{0}",
}

// https://github.com/go-playground/validator/blob/v10.14.0/translations/zh/zh.go
var messageItemsZh = MessageItems{
	{
		tag:         "required",
		translation: "{0}为必填字段",
		override:    false,
	},
	{
		tag:         "required_if",
		translation: "{0}为必填字段",
		override:    false,
	},
	{
		tag:         "required_unless",
		translation: "{0}为必填字段",
		override:    false,
	},
	{
		tag:         "required_with",
		translation: "{0}为必填字段",
		override:    false,
	},
	{
		tag:         "required_with_all",
		translation: "{0}为必填字段",
		override:    false,
	},
	{
		tag:         "required_without",
		translation: "{0}为必填字段",
		override:    false,
	},
	{
		tag:         "required_without_all",
		translation: "{0}为必填字段",
		override:    false,
	},
	{
		tag: "len",
		customRegisFunc: func() (items MessageItems) {

			items.Add("len-string", "{0}长度必须是{1}", false)
			items.Add("len-number", "{0}必须等于{1}", false)
			items.Add("len-items", "{0}必须包含{1}", false)

			return
		},
		customTransFunc: func(fe validator.FieldError, items MessageItems, fmtCardinal FmtCardinal) (t string) {
			var err error
			var kind reflect.Kind
			field := fe.Field()
			param := fe.Param()

			kind = fe.Kind()
			if kind == reflect.Ptr {
				kind = fe.Type().Elem().Kind()
			}

			switch kind {
			case reflect.String:

				oFind, ok := items.Find("len-string")
				if !ok {
					goto END
				}
				c := fmtCardinal("character", param)
				t = oFind.T(field, c)
				// c, err = ut.C("len-string-character", f64, digits, ut.FmtNumber(f64, digits))
				// t, err = ut.T("len-string", fe.Field(), c)

			case reflect.Slice, reflect.Map, reflect.Array:

				oFind, ok := items.Find("len-items")
				if !ok {
					goto END
				}
				c := fmtCardinal("items-item", param)
				t = oFind.T(field, c)
				// c, err = ut.C("len-items-item", f64, digits, ut.FmtNumber(f64, digits))
				// t, err = ut.T("len-items", fe.Field(), c)

			default:

				oFind, ok := items.Find("len-number")
				if !ok {
					goto END
				}
				c := fmtCardinal("placeholder", param)
				t = oFind.T(field, c)
				// t, err = ut.T("len-number", fe.Field(), ut.FmtNumber(f64, digits))

			}

		END:
			if err != nil {
				fmt.Printf("警告: 翻译字段错误: %s", err)
				return fe.(error).Error()
			}

			return t
		},
	},
	{
		tag: "min",
		customRegisFunc: func() (items MessageItems) {

			items.Add("min-string", "{0}长度必须至少为{1}", false)
			items.Add("min-number", "{0}最小只能为{1}", false)
			items.Add("min-items", "{0}必须至少包含{1}", false)

			return
		},
		customTransFunc: func(fe validator.FieldError, items MessageItems, fmtCardinal FmtCardinal) (t string) {
			var err error
			var kind reflect.Kind
			field := fe.Field()
			param := fe.Param()

			kind = fe.Kind()
			if kind == reflect.Ptr {
				kind = fe.Type().Elem().Kind()
			}

			switch kind {
			case reflect.String:

				oFind, ok := items.Find("min-string")
				if !ok {
					goto END
				}
				c := fmtCardinal("character", param)
				t = oFind.T(field, c)
				// c, err = ut.C("min-string-character", f64, digits, ut.FmtNumber(f64, digits))
				// t, err = ut.T("min-string", fe.Field(), c)

			case reflect.Slice, reflect.Map, reflect.Array:

				oFind, ok := items.Find("min-items")
				if !ok {
					goto END
				}
				c := fmtCardinal("items-item", param)
				t = oFind.T(field, c)
				// c, err = ut.C("min-items-item", f64, digits, ut.FmtNumber(f64, digits))
				// t, err = ut.T("min-items", fe.Field(), c)

			default:

				oFind, ok := items.Find("min-number")
				if !ok {
					goto END
				}
				c := fmtCardinal("placeholder", param)
				t = oFind.T(field, c)
				// t, err = ut.T("min-number", fe.Field(), ut.FmtNumber(f64, digits))

			}

		END:
			if err != nil {
				fmt.Printf("警告: 翻译字段错误: %s", err)
				return fe.(error).Error()
			}

			return t
		},
	},
	{
		tag: "max",
		customRegisFunc: func() (items MessageItems) {

			items.Add("max-string", "{0}长度不能超过{1}", false)
			items.Add("max-number", "{0}必须小于或等于{1}", false)
			items.Add("max-items", "{0}最多只能包含{1}", false)

			return
		},
		customTransFunc: func(fe validator.FieldError, items MessageItems, fmtCardinal FmtCardinal) (t string) {
			var err error
			var kind reflect.Kind
			field := fe.Field()
			param := fe.Param()

			kind = fe.Kind()
			if kind == reflect.Ptr {
				kind = fe.Type().Elem().Kind()
			}

			switch kind {
			case reflect.String:

				oFind, ok := items.Find("max-string")
				if !ok {
					goto END
				}
				c := fmtCardinal("character", param)
				t = oFind.T(field, c)
				// c, err = ut.C("max-string-character", f64, digits, ut.FmtNumber(f64, digits))
				// t, err = ut.T("max-string", fe.Field(), c)

			case reflect.Slice, reflect.Map, reflect.Array:

				oFind, ok := items.Find("max-items")
				if !ok {
					goto END
				}
				c := fmtCardinal("items-item", param)
				t = oFind.T(field, c)
				// c, err = ut.C("max-items-item", f64, digits, ut.FmtNumber(f64, digits))
				// t, err = ut.T("max-items", fe.Field(), c)

			default:

				oFind, ok := items.Find("max-number")
				if !ok {
					goto END
				}
				c := fmtCardinal("placeholder", param)
				t = oFind.T(field, c)
				// t, err = ut.T("max-number", fe.Field(), ut.FmtNumber(f64, digits))

			}

		END:
			if err != nil {
				fmt.Printf("警告: 翻译字段错误: %s", err)
				return fe.(error).Error()
			}

			return t
		},
	},
	{
		tag:         "eq",
		translation: "{0}不等于{1}",
		override:    false,
	},
	{
		tag:         "ne",
		translation: "{0}不能等于{1}",
		override:    false,
	},
	{
		tag: "lt",
		customRegisFunc: func() (items MessageItems) {

			items.Add("lt-string", "{0}长度必须小于{1}", false)
			items.Add("lt-number", "{0}必须小于{1}", false)
			items.Add("lt-items", "{0}必须包含少于{1}", false)
			items.Add("lt-datetime", "{0}必须小于当前日期和时间", false)

			return
		},
		customTransFunc: func(fe validator.FieldError, items MessageItems, fmtCardinal FmtCardinal) (t string) {
			var err error
			var kind reflect.Kind
			field := fe.Field()
			param := fe.Param()

			kind = fe.Kind()
			if kind == reflect.Ptr {
				kind = fe.Type().Elem().Kind()
			}

			switch kind {
			case reflect.String:

				oFind, ok := items.Find("lt-string")
				if !ok {
					goto END
				}
				c := fmtCardinal("character", param)
				t = oFind.T(field, c)
				// c, err = ut.C("lt-string-character", f64, digits, ut.FmtNumber(f64, digits))
				// t, err = ut.T("lt-string", fe.Field(), c)

			case reflect.Slice, reflect.Map, reflect.Array:

				oFind, ok := items.Find("lt-items")
				if !ok {
					goto END
				}
				c := fmtCardinal("items-item", param)
				t = oFind.T(field, c)
			// c, err = ut.C("lt-items-item", f64, digits, ut.FmtNumber(f64, digits))
			// t, err = ut.T("lt-items", fe.Field(), c)

			case reflect.Struct:
				if fe.Type() != reflect.TypeOf(time.Time{}) {
					err = fmt.Errorf("tag '%s'不能用于struct类型", fe.Tag())
					goto END
				}

				oFind, ok := items.Find("lt-datetime")
				if !ok {
					goto END
				}
				t = oFind.T(field)
				// t, err = ut.T("lt-datetime", fe.Field())

			default:

				oFind, ok := items.Find("lt-number")
				if !ok {
					goto END
				}
				c := fmtCardinal("placeholder", param)
				t = oFind.T(field, c)
				// t, err = ut.T("lt-number", fe.Field(), ut.FmtNumber(f64, digits))

			}

		END:
			if err != nil {
				fmt.Printf("警告: 翻译字段错误: %s", err)
				return fe.(error).Error()
			}

			return t
		},
	},
	{
		tag: "lte",
		customRegisFunc: func() (items MessageItems) {

			items.Add("lte-string", "{0}长度不能超过{1}", false)
			items.Add("lte-number", "{0}必须小于或等于{1}", false)
			items.Add("lte-items", "{0}最多只能包含{1}", false)
			items.Add("lte-datetime", "{0}必须小于或等于当前日期和时间", false)

			return
		},
		customTransFunc: func(fe validator.FieldError, items MessageItems, fmtCardinal FmtCardinal) (t string) {
			var err error
			var kind reflect.Kind
			field := fe.Field()
			param := fe.Param()

			kind = fe.Kind()
			if kind == reflect.Ptr {
				kind = fe.Type().Elem().Kind()
			}

			switch kind {
			case reflect.String:

				oFind, ok := items.Find("lte-string")
				if !ok {
					goto END
				}
				c := fmtCardinal("character", param)
				t = oFind.T(field, c)
				// c, err = ut.C("lte-string-character", f64, digits, ut.FmtNumber(f64, digits))
				// t, err = ut.T("lte-string", fe.Field(), c)

			case reflect.Slice, reflect.Map, reflect.Array:

				oFind, ok := items.Find("lte-items")
				if !ok {
					goto END
				}
				c := fmtCardinal("items-item", param)
				t = oFind.T(field, c)
			// c, err = ut.C("lte-items-item", f64, digits, ut.FmtNumber(f64, digits))
			// t, err = ut.T("lte-items", fe.Field(), c)

			case reflect.Struct:
				if fe.Type() != reflect.TypeOf(time.Time{}) {
					err = fmt.Errorf("tag '%s'不能用于struct类型", fe.Tag())
					goto END
				}

				oFind, ok := items.Find("lte-datetime")
				if !ok {
					goto END
				}
				t = oFind.T(field)
				// t, err = ut.T("lte-datetime", fe.Field())

			default:

				oFind, ok := items.Find("lte-number")
				if !ok {
					goto END
				}
				c := fmtCardinal("placeholder", param)
				t = oFind.T(field, c)
				// t, err = ut.T("lte-number", fe.Field(), ut.FmtNumber(f64, digits))

			}

		END:
			if err != nil {
				fmt.Printf("警告: 翻译字段错误: %s", err)
				return fe.(error).Error()
			}

			return t
		},
	},
	{
		tag: "gt",
		customRegisFunc: func() (items MessageItems) {

			items.Add("gt-string", "{0}长度必须大于{1}", false)
			items.Add("gt-number", "{0}必须大于{1}", false)
			items.Add("gt-items", "{0}必须大于{1}", false)
			items.Add("gt-datetime", "{0}必须大于当前日期和时间", false)

			return
		},
		customTransFunc: func(fe validator.FieldError, items MessageItems, fmtCardinal FmtCardinal) (t string) {
			var err error
			var kind reflect.Kind
			field := fe.Field()
			param := fe.Param()

			kind = fe.Kind()
			if kind == reflect.Ptr {
				kind = fe.Type().Elem().Kind()
			}

			switch kind {
			case reflect.String:

				oFind, ok := items.Find("gt-string")
				if !ok {
					goto END
				}
				c := fmtCardinal("character", param)
				t = oFind.T(field, c)
				// c, err = ut.C("gt-string-character", f64, digits, ut.FmtNumber(f64, digits))
				// t, err = ut.T("gt-string", fe.Field(), c)

			case reflect.Slice, reflect.Map, reflect.Array:

				oFind, ok := items.Find("gt-items")
				if !ok {
					goto END
				}
				c := fmtCardinal("items-item", param)
				t = oFind.T(field, c)
			// c, err = ut.C("gt-items-item", f64, digits, ut.FmtNumber(f64, digits))
			// t, err = ut.T("gt-items", fe.Field(), c)

			case reflect.Struct:
				if fe.Type() != reflect.TypeOf(time.Time{}) {
					err = fmt.Errorf("tag '%s'不能用于struct类型", fe.Tag())
					goto END
				}

				oFind, ok := items.Find("gt-datetime")
				if !ok {
					goto END
				}
				t = oFind.T(field)
				// t, err = ut.T("gt-datetime", fe.Field())

			default:

				oFind, ok := items.Find("gt-number")
				if !ok {
					goto END
				}
				c := fmtCardinal("placeholder", param)
				t = oFind.T(field, c)
				// t, err = ut.T("gt-number", fe.Field(), ut.FmtNumber(f64, digits))

			}

		END:
			if err != nil {
				fmt.Printf("警告: 翻译字段错误: %s", err)
				return fe.(error).Error()
			}

			return t
		},
	},
	{
		tag: "gte",
		customRegisFunc: func() (items MessageItems) {

			items.Add("gte-string", "{0}长度必须至少为{1}", false)
			items.Add("gte-number", "{0}必须大于或等于{1}", false)
			items.Add("gte-items", "{0}必须至少包含{1}", false)
			items.Add("gte-datetime", "{0}必须大于或等于当前日期和时间", false)

			return
		},
		customTransFunc: func(fe validator.FieldError, items MessageItems, fmtCardinal FmtCardinal) (t string) {
			var err error
			var kind reflect.Kind
			field := fe.Field()
			param := fe.Param()

			kind = fe.Kind()
			if kind == reflect.Ptr {
				kind = fe.Type().Elem().Kind()
			}

			switch kind {
			case reflect.String:

				oFind, ok := items.Find("gte-string")
				if !ok {
					goto END
				}
				c := fmtCardinal("character", param)
				t = oFind.T(field, c)
				// c, err = ut.C("gte-string-character", f64, digits, ut.FmtNumber(f64, digits))
				// t, err = ut.T("gte-string", fe.Field(), c)

			case reflect.Slice, reflect.Map, reflect.Array:

				oFind, ok := items.Find("gte-items")
				if !ok {
					goto END
				}
				c := fmtCardinal("items-item", param)
				t = oFind.T(field, c)
			// c, err = ut.C("gte-items-item", f64, digits, ut.FmtNumber(f64, digits))
			// t, err = ut.T("gte-items", fe.Field(), c)

			case reflect.Struct:
				if fe.Type() != reflect.TypeOf(time.Time{}) {
					err = fmt.Errorf("tag '%s'不能用于struct类型", fe.Tag())
					goto END
				}

				oFind, ok := items.Find("gte-datetime")
				if !ok {
					goto END
				}
				t = oFind.T(field)
				// t, err = ut.T("gte-datetime", fe.Field())

			default:

				oFind, ok := items.Find("gte-number")
				if !ok {
					goto END
				}
				c := fmtCardinal("placeholder", param)
				t = oFind.T(field, c)
				// t, err = ut.T("gte-number", fe.Field(), ut.FmtNumber(f64, digits))

			}

		END:
			if err != nil {
				fmt.Printf("警告: 翻译字段错误: %s", err)
				return fe.(error).Error()
			}

			return t
		},
	},
	{
		tag:         "eqfield",
		translation: "{0}必须等于{1}",
		override:    false,
	},
	{
		tag:         "eqcsfield",
		translation: "{0}必须等于{1}",
		override:    false,
	},
	{
		tag:         "necsfield",
		translation: "{0}不能等于{1}",
		override:    false,
	},
	{
		tag:         "gtcsfield",
		translation: "{0}必须大于{1}",
		override:    false,
	},
	{
		tag:         "gtecsfield",
		translation: "{0}必须大于或等于{1}",
		override:    false,
	},
	{
		tag:         "ltcsfield",
		translation: "{0}必须小于{1}",
		override:    false,
	},
	{
		tag:         "ltecsfield",
		translation: "{0}必须小于或等于{1}",
		override:    false,
	},
	{
		tag:         "nefield",
		translation: "{0}不能等于{1}",
		override:    false,
	},
	{
		tag:         "gtfield",
		translation: "{0}必须大于{1}",
		override:    false,
	},
	{
		tag:         "gtefield",
		translation: "{0}必须大于或等于{1}",
		override:    false,
	},
	{
		tag:         "ltfield",
		translation: "{0}必须小于{1}",
		override:    false,
	},
	{
		tag:         "ltefield",
		translation: "{0}必须小于或等于{1}",
		override:    false,
	},
	{
		tag:         "alpha",
		translation: "{0}只能包含字母",
		override:    false,
	},
	{
		tag:         "alphanum",
		translation: "{0}只能包含字母和数字",
		override:    false,
	},
	{
		tag:         "alphanumunicode",
		translation: "{0}只能包含字母数字和Unicode字符",
		override:    false,
	},
	{
		tag:         "alphaunicode",
		translation: "{0}只能包含字母和Unicode字符",
		override:    false,
	},
	{
		tag:         "numeric",
		translation: "{0}必须是一个有效的数值",
		override:    false,
	},
	{
		tag:         "number",
		translation: "{0}必须是一个有效的数字",
		override:    false,
	},
	{
		tag:         "hexadecimal",
		translation: "{0}必须是一个有效的十六进制",
		override:    false,
	},
	{
		tag:         "hexcolor",
		translation: "{0}必须是一个有效的十六进制颜色",
		override:    false,
	},
	{
		tag:         "rgb",
		translation: "{0}必须是一个有效的RGB颜色",
		override:    false,
	},
	{
		tag:         "rgba",
		translation: "{0}必须是一个有效的RGBA颜色",
		override:    false,
	},
	{
		tag:         "hsl",
		translation: "{0}必须是一个有效的HSL颜色",
		override:    false,
	},
	{
		tag:         "hsla",
		translation: "{0}必须是一个有效的HSLA颜色",
		override:    false,
	},
	{
		tag:         "email",
		translation: "{0}必须是一个有效的邮箱",
		override:    false,
	},
	{
		tag:         "url",
		translation: "{0}必须是一个有效的URL",
		override:    false,
	},
	{
		tag:         "uri",
		translation: "{0}必须是一个有效的URI",
		override:    false,
	},
	{
		tag:         "base64",
		translation: "{0}必须是一个有效的Base64字符串",
		override:    false,
	},
	{
		tag:         "contains",
		translation: "{0}必须包含文本'{1}'",
		override:    false,
	},
	{
		tag:         "containsany",
		translation: "{0}必须包含至少一个以下字符'{1}'",
		override:    false,
	},
	{
		tag:         "containsrune",
		translation: "{0}必须包含字符'{1}'",
		override:    false,
	},

	{
		tag:         "excludes",
		translation: "{0}不能包含文本'{1}'",
		override:    false,
	},
	{
		tag:         "excludesall",
		translation: "{0}不能包含以下任何字符'{1}'",
		override:    false,
	},
	{
		tag:         "excludesrune",
		translation: "{0}不能包含'{1}'",
		override:    false,
	},
	{
		tag:         "endswith",
		translation: "{0}必须以文本'{1}'结尾",
		override:    false,
	},
	{
		tag:         "startswith",
		translation: "{0}必须以文本'{1}'开头",
		override:    false,
	},
	{
		tag:         "isbn",
		translation: "{0}必须是一个有效的ISBN编号",
		override:    false,
	},
	{
		tag:         "isbn10",
		translation: "{0}必须是一个有效的ISBN-10编号",
		override:    false,
	},
	{
		tag:         "isbn13",
		translation: "{0}必须是一个有效的ISBN-13编号",
		override:    false,
	},
	{
		tag:         "uuid",
		translation: "{0}必须是一个有效的UUID",
		override:    false,
	},
	{
		tag:         "uuid3",
		translation: "{0}必须是一个有效的V3 UUID",
		override:    false,
	},
	{
		tag:         "uuid4",
		translation: "{0}必须是一个有效的V4 UUID",
		override:    false,
	},
	{
		tag:         "uuid5",
		translation: "{0}必须是一个有效的V5 UUID",
		override:    false,
	},
	{
		tag:         "ulid",
		translation: "{0}必须是一个有效的ULID",
		override:    false,
	},
	{
		tag:         "ascii",
		translation: "{0}必须只包含ascii字符",
		override:    false,
	},
	{
		tag:         "printascii",
		translation: "{0}必须只包含可打印的ascii字符",
		override:    false,
	},
	{
		tag:         "multibyte",
		translation: "{0}必须包含多字节字符",
		override:    false,
	},
	{
		tag:         "datauri",
		translation: "{0}必须包含有效的数据URI",
		override:    false,
	},
	{
		tag:         "latitude",
		translation: "{0}必须包含有效的纬度坐标",
		override:    false,
	},
	{
		tag:         "longitude",
		translation: "{0}必须包含有效的经度坐标",
		override:    false,
	},
	{
		tag:         "ssn",
		translation: "{0}必须是一个有效的社会安全号码(SSN)",
		override:    false,
	},
	{
		tag:         "ipv4",
		translation: "{0}必须是一个有效的IPv4地址",
		override:    false,
	},
	{
		tag:         "ipv6",
		translation: "{0}必须是一个有效的IPv6地址",
		override:    false,
	},
	{
		tag:         "ip",
		translation: "{0}必须是一个有效的IP地址",
		override:    false,
	},
	{
		tag:         "cidr",
		translation: "{0}必须是一个有效的无类别域间路由(CIDR)",
		override:    false,
	},
	{
		tag:         "cidrv4",
		translation: "{0}必须是一个包含IPv4地址的有效无类别域间路由(CIDR)",
		override:    false,
	},
	{
		tag:         "cidrv6",
		translation: "{0}必须是一个包含IPv6地址的有效无类别域间路由(CIDR)",
		override:    false,
	},
	{
		tag:         "tcp_addr",
		translation: "{0}必须是一个有效的TCP地址",
		override:    false,
	},
	{
		tag:         "tcp4_addr",
		translation: "{0}必须是一个有效的IPv4 TCP地址",
		override:    false,
	},
	{
		tag:         "tcp6_addr",
		translation: "{0}必须是一个有效的IPv6 TCP地址",
		override:    false,
	},
	{
		tag:         "udp_addr",
		translation: "{0}必须是一个有效的UDP地址",
		override:    false,
	},
	{
		tag:         "udp4_addr",
		translation: "{0}必须是一个有效的IPv4 UDP地址",
		override:    false,
	},
	{
		tag:         "udp6_addr",
		translation: "{0}必须是一个有效的IPv6 UDP地址",
		override:    false,
	},
	{
		tag:         "ip_addr",
		translation: "{0}必须是一个有效的IP地址",
		override:    false,
	},
	{
		tag:         "ip4_addr",
		translation: "{0}必须是一个有效的IPv4地址",
		override:    false,
	},
	{
		tag:         "ip6_addr",
		translation: "{0}必须是一个有效的IPv6地址",
		override:    false,
	},
	{
		tag:         "unix_addr",
		translation: "{0}必须是一个有效的UNIX地址",
		override:    false,
	},
	{
		tag:         "mac",
		translation: "{0}必须是一个有效的MAC地址",
		override:    false,
	},
	{
		tag:         "iscolor",
		translation: "{0}必须是一个有效的颜色",
		override:    false,
	},
	{
		tag:         "oneof",
		translation: "{0}必须是[{1}]中的一个",
		override:    false,
	},

	{
		tag:         "json",
		translation: "{0}必须是一个JSON字符串",
		override:    false,
	},
	{
		tag:         "lowercase",
		translation: "{0}必须是小写字母",
		override:    false,
	},
	{
		tag:         "uppercase",
		translation: "{0}必须是大写字母",
		override:    false,
	},
	{
		tag:         "datetime",
		translation: "{0}的格式必须是{1}",
		override:    false,
	},
	{
		tag:         "image",
		translation: "{0} 必须是有效图像",
		override:    false,
	},

	{
		tag:         "e164",
		translation: "{0}必须是有效的E.164格式的电话号码",
		override:    false,
	},
	{
		tag:         "fqdn",
		translation: "{0}必须是有效的FQDN",
		override:    false,
	},
	{
		tag:         "unique",
		translation: "{0}必须包含唯一值",
		override:    false,
	},
	{
		tag:         "cron",
		translation: "{0}必须是有效的cron表达式",
		override:    false,
	},
	{
		tag:         "jwt",
		translation: "{0}必须是有效的jwt字符串",
		override:    false,
	},
	{
		tag:         "postcode_iso3166_alpha2",
		translation: "{0}与{1}(国家/地区)的邮政编码格式不匹配",
		override:    false,
	},
	{
		tag:         "postcode_iso3166_alpha2_field",
		translation: "{0}(国家/地区)与{1}(国家/地区)的邮政编码格式不匹配",
		override:    false,
	},
	{
		tag:         "boolean",
		translation: "{0}必须是有效的布尔值",
		override:    false,
	},
	{
		tag:         "cve",
		translation: "{0}必须是有效的cve标识符",
		override:    false,
	},
}
