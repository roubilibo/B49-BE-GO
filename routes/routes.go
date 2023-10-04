package routes

import "github.com/labstack/echo/v4"

func RouteInit(e *echo.Group) {
	PaslonRoutes(e)
	UserRoutes(e)
}