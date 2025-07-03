package bootstrap

import (
	"fmt"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"gorm.io/gorm"
	"log"
	"os"
)

func NewDatabase(env *Env) *gorm.DB {
	switch env.DBType {
	case "mysql":
		return GormMysql(env)
	case "pgsql":
		return GormPgSql(env)
	default:
		return GormMysql(env)
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables(db *gorm.DB) {
	// db.Set("gorm:table_options", "CHARSET=utf8mb4").AutoMigrate
	fmt.Println("register tables ...")
	err := db.AutoMigrate(
		domain.User{},
		domain.Task{},
	)
	fmt.Println("register tables success")
	if err != nil {
		os.Exit(0)
	}
}

func CloseDBConnection(db *gorm.DB) {
	if db == nil {
		return
	}

	// 程序结束前关闭数据库链接
	DB, _ := db.DB()
	err := DB.Close()
	if err != nil {
		return
	}
	log.Println("Connection to Database closed.")
}
