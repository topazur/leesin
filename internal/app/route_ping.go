package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (a *App) AddRouteForPing() {

	// curl -X GET http://localhost:9092/ping
	a.engine.GET("/ping", func(ctx *gin.Context) {
		// Sleep 可帮助验证优雅关机的效果
		time.Sleep(5 * time.Second)
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
