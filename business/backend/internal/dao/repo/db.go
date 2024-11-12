package repo

import (
	"backend/internal/config"
	"common/db"
)

func NewDB(c *config.Config) (db.DB, func(), error) {
	var d db.DB
	if len(c.Database.Mysql.Host) > 0 {
		d = db.NewMysql(c.Database.Mysql.Host)
	} else {
		d = db.NewMgo(c.Database.Mongodb.Host, c.Database.Mongodb.DBName)
	}
	return d, func() { _ = d.Close() }, nil
}
