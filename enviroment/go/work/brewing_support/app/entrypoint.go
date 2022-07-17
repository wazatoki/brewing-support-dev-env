package app

import (
	"brewing_support/app/infrastructures/echo"
	"brewing_support/app/infrastructures/postgresql"
	"brewing_support/app/infrastructures/viper"
)

// StartApp entoru point of application
func StartApp() {
	// configuration
	viper.SetupAppConfig()
	postgresql.Migrate()
	echo.StartEcho()
}
