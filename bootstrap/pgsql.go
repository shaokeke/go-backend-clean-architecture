package bootstrap

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormPgSql() *gorm.DB {

	p := ENV.Pgsql
	if p.Dbname == "" {
		return nil
	}
	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: p.Dsn(),

		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {

		log.Println(p.DriverName, err)
	}
	return db
}
