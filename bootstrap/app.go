package bootstrap

import (
	"gorm.io/gorm"
)

type Application struct {
	Env   *Env
	OrmDb *gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.OrmDb = NewDatabase(app.Env)
	return *app
}
func (app *Application) RegisterTables() {
	RegisterTables(app.OrmDb)
}

func (app *Application) CloseDBConnection() {
	CloseDBConnection(app.OrmDb)
}
