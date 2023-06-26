package app

import (
	"github.com/gin-gonic/gin"
	"github.com/topazur/leesin/internal/controller"
	"github.com/topazur/leesin/pkg/log"
	"github.com/topazur/leesin/pkg/token"
)

type App struct {
	engine     *gin.Engine
	logger     *log.Logger
	tokenMaker token.Maker
	handler    *controller.Controller
}

func NewApp(logger *log.Logger, tokenMaker token.Maker, handler *controller.Controller) *App {
	gin.SetMode(gin.ReleaseMode)

	/* 创建 Engine 实例 */
	// engine := gin.New()
	engine := gin.Default()

	return &App{
		engine:     engine,
		logger:     logger,
		tokenMaker: tokenMaker,
		handler:    handler,
	}
}
