package routes

import (
	"myapp/handlers"

	"github.com/labstack/echo/v4"
)

func PaslonRoutes(e *echo.Group) {
	e.GET("/paslons", handlers.FindPaslons)
	e.GET("/paslon/:id", handlers.GetPaslon)
	e.POST("/paslon", handlers.CreatePaslon)
	e.PATCH("/paslon/:id", handlers.UpdatePaslon)
	e.DELETE("/paslon/:id", handlers.DeletePaslon)
}
