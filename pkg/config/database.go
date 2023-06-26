package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

// NewDBConnByPgx 通过 pgx/v5 连接数据库
func NewDBConnByPgx(connString string) (*pgx.Conn, error) {
	// 1️⃣ pgx.ParseConfig 函数用于解析 PostgreSQL 连接字符串，返回一个包含连接参数的 pgx.ConnConfig 对象
	// config, err := pgx.ParseConfig(connString)
	// if err != nil {
	// 	log.Fatalf("[Connect] pgx.Connect error: %v", err)
	// 	os.Exit(1)
	// 	return nil, nil
	// }
	// conn, err := pgx.ConnectConfig(context.Background(), config)

	// 2️⃣ 连接数据库 (可以直接传递连接字符串，也可以传递 pgx.ConnConfig 对象)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("[Connect] pgx.Connect error: %v", err)
	}

	// NOTICE: 不能在此处关闭，否则该函数调用结束就会关闭连接；应该在业务调用方处理关闭。
	// defer conn.Close(context.Background())
	return conn, nil
}
