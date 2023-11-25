package routes

import (
	"miniproject/handler"
	"miniproject/middleware"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(e *echo.Echo) {

	e.POST("/users/register", handler.RegisterUser)
	e.POST("/users/login", handler.LoginUser)

	e.POST("/users/profile/create", handler.CreateUserProfile, middleware.RequireAuth)
	e.GET("/users/profile", handler.GetUserProfile, middleware.RequireAuth)
	e.PUT("/users/profile/update", handler.UpdateUserProfile, middleware.RequireAuth)

}
