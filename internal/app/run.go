package app

import "github.com/topazur/leesin/pkg/http"

// Run 启动服务(实现了优雅关机)
func (a *App) Run(addr string) {
	http.RunAtGracefulShutdown(a.engine, addr)
}
