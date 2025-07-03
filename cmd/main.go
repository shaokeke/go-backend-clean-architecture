package main

import (
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/api/route"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {

	app := bootstrap.App()

	env := app.Env

	db := app.OrmDb
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	ginEngine := gin.Default()

	route.Setup(env, timeout, db, ginEngine)

	ginEngine.Run(env.ServerAddress)

}
