package bootstrap

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GormMysql(env *Env) *gorm.DB {

	if env.DBName == "" {
		return nil
	}
	db, err := gorm.Open(mysql.Open(env.mysqlDsn()), &gorm.Config{})
	if err != nil {
		log.Println(env.DBType, err)
	}
	return db

}
