package db

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	client *gorm.DB
}

// NewMysql init db
func NewMysql(url string) *Mysql {
	m := new(Mysql)
	err := m.initDB(url)
	if err != nil {
		panic(err)
	}
	return m
}

func (m *Mysql) initDB(url string) (err error) {
	m.client, err = gorm.Open(mysql.Open(url), &gorm.Config{})
	return
}

func (m Mysql) Save(ctx context.Context, tb string, item any) (err error) {
	return m.client.WithContext(ctx).Table(tb).Create(item).Error
}

func (m Mysql) Find(ctx context.Context, tb string, filter map[string]any, item any) (err error) {
	m.client.WithContext(ctx).Table(tb).Where(filter).Find(item)
	return
}

func (m Mysql) List(
	ctx context.Context,
	tb string,
	filter map[string]any,
	orderBy string,
	page, pageSize int,
	list []any,
) (err error) {
	limit := pageSize
	offset := (page - 1) * pageSize
	err = m.client.
		WithContext(ctx).
		Table(tb).
		Where(filter).
		Order(orderBy).
		Limit(limit).
		Offset(offset).
		Find(&list).Error
	return
}

func (m Mysql) Delete(ctx context.Context, tb string, filter map[string]any) (err error) {
	return m.client.WithContext(ctx).Table(tb).Where(filter).Delete(nil).Error
}

func (m Mysql) Close() error {
	db, err := m.client.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
