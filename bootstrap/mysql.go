package bootstrap

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql() *gorm.DB {

	m := ENV.Mysql
	if m.Dbname == "" {
		return nil
	}
	db, err := gorm.Open(mysql.Open(m.Dsn()), &gorm.Config{})
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(m.DriverName, err)
	}
	return db

}
