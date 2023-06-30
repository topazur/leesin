package config

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// NewDBConnByPgx 通过 pgx/v5 连接数据库
func NewDBConnByPgx(connString string) (*pgx.Conn, error) {
	/// 1️⃣ pgx.ParseConfig 函数用于解析 PostgreSQL 连接字符串，返回一个包含连接参数的 pgx.ConnConfig 对象
	// config, err := pgx.ParseConfig(connString)
	// if err != nil {
	// 	return nil, err
	// }
	// conn, err := pgx.ConnectConfig(context.Background(), config)

	/// 2️⃣ 连接数据库 (可以直接传递连接字符串，也可以传递 pgx.ConnConfig 对象)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("[Connect] pgx.Connect error: %v", err)
	}

	// NOTICE: 不能在此处关闭，否则该函数调用结束就会关闭连接；应该在业务调用方处理关闭。
	// defer conn.Close(context.Background())
	return conn, nil
}

// NewDBConnByPgx 通过 pgx/v5/pgxpool 连接数据库
// pgxpool是用于pgx的一个并发安全的连接池。
// pgxpool实现了与pgx连接几乎相同的接口。
func NewDBConnByPgxpool(connString string) (*pgxpool.Pool, error) {
	/// 1️⃣ pgx.ParseConfig => 可指定 PostgreSQL设置、pgx设置和池设置
	// config, err := pgxpool.ParseConfig(connString)
	// if err != nil {
	// 	return nil, err
	// }
	// config.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {}
	// connPool, err := pgxpool.NewWithConfig(context.Background(), config)

	/// 2️⃣ 连接数据库 (可以直接传递连接字符串，也可以传递 pgx.ConnConfig 对象)
	connPool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("[Connect] pgx.Connect error: %v", err)
	}

	// NOTICE: 不能在此处关闭，否则该函数调用结束就会关闭连接；应该在业务调用方处理关闭。
	// defer conn.Close(context.Background())
	return connPool, nil
}
