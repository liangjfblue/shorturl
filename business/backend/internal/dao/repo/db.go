package repo

import (
	"backend/internal/config"
	"common/db"
)

func NewDB(c *config.Config) db.DB {
	if len(c.Database.Mysql.Host) > 0 {
		return db.NewMysql(c.Database.Mysql.Host)
	}
	return db.NewMgo(c.Database.Mongo.Host, c.Database.Mongo.DBName)
}
