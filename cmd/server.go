package cmd

import (
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/topazur/leesin/di"
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

		/// 获取配置文件
		conf := config.NewConfig(configPath)
		/// 实例化zap对象
		logger := log.NewLog(conf)

		/// wire 依赖注入
		app, closeup, err := di.InitializeGinServer(conf, logger)
		defer closeup()
		if err != nil {
			cmd.Println(err)
			return
		}

		addr := config.GetVariableString(conf, conf.GetString("http.gin_addr"))
		logger.Info("server start", zap.String("host", "http://"+addr))

		/// 启动服务
		app.Run(addr)
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
