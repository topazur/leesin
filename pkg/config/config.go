package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// NewConfig 通过viper读取配置文件，可手动传入配置文件路径
// 放弃维护结构体解码(viper.Unmarshal)的方式，直接使用viper.GetXxx()获取配置
// [解码自定义格式 - Unmarshal和UnmarshalKey](https://sagikazarmark.hu/blog/decoding-custom-formats-with-viper/)
func NewConfig(path string) *viper.Viper {
	return initViper(path)
}

func initViper(path string) (conf *viper.Viper) {
	conf = viper.New()

	absPath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("[Config] filepath.Abs error: %v", err)
		os.Exit(1)
		return
	}

	// AddConfigPath + SetConfigName + SetConfigType
	conf.SetConfigFile(absPath)

	// 搜索存在的第一个路径(并且是一个配置文件); 如果没有找到，则返回错误。
	if err := conf.ReadInConfig(); err != nil {
		log.Fatalf("[Config] ReadInConfig error: %v", err)
		os.Exit(1)
		return
	}

	return
}
