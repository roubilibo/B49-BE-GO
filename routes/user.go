package routes

import (
	"myapp/handlers"
	"myapp/pkg/mysql"
	"myapp/repositories"

	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Group) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	e.GET("/users", h.FindUsers)
	e.GET("/user/:id", h.GetUser)
	e.POST("/users", h.AddUser)
	e.DELETE("/user/:id", h.DeleteUser)
	e.PATCH("/user/:id", h.UpdateUser)
}
