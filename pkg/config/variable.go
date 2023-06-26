package config

import (
	"regexp"

	"github.com/spf13/viper"
)

// GetVariableString 由于 Viper 默认不会识别 YAML 文件中的变量，所以需要自定义函数来获取变量
// 目前仅支持 ${} 语法的变量
func GetVariableString(conf *viper.Viper, key string) string {
	// 构建正则表达式
	re := regexp.MustCompile(`\$\{([^}]+)\}`)

	// 通过正则表达式匹配出所有的变量(-1 表示匹配所有) => [[${db.postgres.username} db.postgres.username], ...]
	// result := re.FindAllStringSubmatch(key, -1)

	// 匹配并可迭代自定义替换内容
	result := re.ReplaceAllStringFunc(key, func(match string) string {
		// `match => ${db.postgres.username}`; `variable => db.postgres.username`
		variable := match[2 : len(match)-1]

		return conf.GetString(variable)
	})

	return result
}
