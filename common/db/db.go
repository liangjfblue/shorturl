package db

import "context"

// DB 数据库操作
type DB interface {
	Save(ctx context.Context, tb string, item any) error
	Find(ctx context.Context, tb string, filter map[string]any, item any) error
	List(ctx context.Context, tb string, filter map[string]any, orderBy string, page, pageSize int, list []any) (err error)
	Delete(ctx context.Context, tb string, filter map[string]any) error
}
