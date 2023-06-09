package service

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/topazur/leesin/pkg/log"
	"github.com/topazur/leesin/pkg/token"
)

// DBTX querier generated by sqlc <便于兼容pgx.Conn 和 pgx.Tx 均可作为参数传入>
type DBTX interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	SendBatch(context.Context, *pgx.Batch) pgx.BatchResults
}

// Store 聚合所有sqlc生成的Querier接口
type Store interface {
}

// StoreImpl 聚合匿名结构体，实现Store接口
type StoreImpl struct {
	db         DBTX
	logger     *log.Logger
	tokenMaker token.Maker
}

// NewStore 实例化
func NewService(db DBTX, logger *log.Logger, tokenMaker token.Maker) (store Store) {
	return &StoreImpl{
		db:         db,
		logger:     logger,
		tokenMaker: tokenMaker,
	}
}
