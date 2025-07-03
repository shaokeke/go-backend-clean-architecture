package main

import (
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/bootstrap"
)

func main() {

	app := bootstrap.App()
	defer app.CloseDBConnection()
	//注册数据库表
	app.RegisterTables()
}
