package bootstrap

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GormPgSql(env *Env) *gorm.DB {

	if env.DBName == "" {
		return nil
	}
	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: env.pgsqlDsn(),

		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {

		log.Println(env.DBType, err)
	}
	return db
}
