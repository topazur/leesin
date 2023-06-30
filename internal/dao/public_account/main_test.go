package dao

import (
	"os"
	"testing"

	"github.com/topazur/leesin/pkg/config"
)

// testQueries 用于测试的 Querier 实例
var testQueries *Queries

// TestMain 同一目录下的所有测试之前执行，可做一些初始化操作
func TestMain(m *testing.M) {
	configPath := "../../../config/config.yaml"
	conf := config.NewConfig(configPath)

	connString := config.GetVariableString(conf, conf.GetString("db.postgres.conn_url"))
	connPool, err := config.NewDBConnByPgxpool(connString)
	if err != nil {
		os.Exit(1)
	}

	// 实例化Querier
	testQueries = New(connPool)

	// 🌈 运行单元测试函数
	exitCode := m.Run()

	// 退出测试 (退出之前可以完成一些清理操作)
	connPool.Close()
	os.Exit(exitCode)
}
