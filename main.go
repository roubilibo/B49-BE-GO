package main

import (
	"fmt"
	"myapp/database"
	"myapp/pkg/mysql"
	"myapp/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	mysql.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	fmt.Println("server running localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}