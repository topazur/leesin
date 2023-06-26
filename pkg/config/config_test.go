package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Test_NewConfig_Exist 测试配置文件存在的情况(相对路径)
func Test_NewConfig_Exist(t *testing.T) {
	conf := NewConfig("../../config/config.example.yaml")

	require.NotNil(t, conf)
	require.NotEmpty(t, conf)
	require.Equal(t, "local", conf.GetString("env"))
}

// Test_GetVariableString 由于 Viper 默认不会识别 YAML 文件中的变量，所以需要自定义函数来获取变量
func Test_GetVariableString(t *testing.T) {
	conf := NewConfig("../../config/config.example.yaml")

	require.NotNil(t, conf)

	result := GetVariableString(conf, conf.GetString("db.postgres.conn_url"))
	require.NotEmpty(t, result)
}
