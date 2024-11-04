package host

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Dsn    string
	Config *gorm.Config
}

var DB *gorm.DB

func initDB() (err error) {
	DB, err = gorm.Open(mysql.Open(Cfg.Database.Dsn), &gorm.Config{})
	return
}
