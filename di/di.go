// github.com/google/wire v0.5.0 // indirect
// 模拟wire的代码结构，但是依赖注入手动实现
package di

import (
	"context"

	"github.com/spf13/viper"
	"github.com/topazur/leesin/internal/app"
	"github.com/topazur/leesin/internal/controller"
	"github.com/topazur/leesin/internal/service"
	"github.com/topazur/leesin/pkg/config"
	"github.com/topazur/leesin/pkg/log"
	"github.com/topazur/leesin/pkg/token"
)

// InitializeGinServer 初始化gin服务
func InitializeGinServer(conf *viper.Viper, logger *log.Logger) (*app.App, func(), error) {
	/// 实例化tokenMaker
	symmetricKey := conf.GetString("security.token.key")
	tokenMaker, err := token.NewPasetoMaker(symmetricKey)
	if err != nil {
		return nil, nil, err
	}

	/// 链接数据库
	connString := config.GetVariableString(conf, conf.GetString("db.postgres.conn_url"))
	conn, err := config.NewDBConnByPgx(connString)
	if err != nil {
		return nil, nil, err
	}

	/// dao and service 层
	store := service.NewService(conn, logger, tokenMaker)
	/// controller 层
	handler := controller.NewController(logger, store)

	/// 实例化Gin引擎
	instance := app.NewApp(logger, tokenMaker, handler)

	/// 注册路由
	instance.AddRouteForPing()

	return instance, func() {
		conn.Close(context.Background())
		logger.Sync()
	}, nil
}
