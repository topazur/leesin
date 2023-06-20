package cmd

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/topazur/leesin/pkg/config"
	"github.com/topazur/leesin/pkg/log"
	"go.uber.org/zap"
)

var configArg = &customFlag{
	name:      "config",
	shorthand: "c",
	usage:     `config path, eg: config/config.yaml`,
}

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Run Service",
	Long:  `Run Service`,
	// NOTICE: [参考该项目注册命令，实现Run方法后才会在help中显示注册的此命令](https://github.com/Jeeejeets/gmail_terminal/blob/master/cmd/send.go)
	Run: func(cmd *cobra.Command, args []string) {
		// cmd.Flags().GetString(configArg.name)
		configPath := cast.ToString(configArg.value)

		conf := config.NewConfig(configPath)
		logger := log.NewLog(conf)

		logger.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))
	},
}

func init() {
	// serverCmd 解析 config 命令行参数 (手动绑定变量 - 赋值接受返回值)
	configArg.value = serverCmd.Flags().StringP(
		configArg.name,
		configArg.shorthand,
		"config/config.yaml",
		configArg.usage,
	)

	rootCmd.AddCommand(serverCmd)
}
