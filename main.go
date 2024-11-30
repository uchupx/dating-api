package main

import (
	"github.com/labstack/echo/v4"
	"github.com/uchupx/dating-api/config"
	dating "github.com/uchupx/dating-api/src"
)

func main() {
	e := echo.New()
	e.Debug = true

	conf := config.GetConfig()
	// initial logger
	d := dating.Dating{}
	runAPIServer(conf, e, &d)
}

func runAPIServer(conf *config.Config, e *echo.Echo, i *dating.Dating) {
	i.InitRoutes(conf, e)
	if err := e.Start(":" + conf.App.Port); err != nil {
		e.Logger.Fatal(err.Error())
	}
}
