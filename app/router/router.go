package router

import (
	"net/http"

	"immersive-dash-4/app/middlewares"
	_userData "immersive-dash-4/features/user/data"
	_userHandler "immersive-dash-4/features/user/handler"
	_userService "immersive-dash-4/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Welcome to Immersive Dash 4",
		})
	})

	//Authentikasi
	e.POST("/login", userHandlerAPI.Login)

	//Users
	e.GET("/users/:user_id", userHandlerAPI.GetUserById, middlewares.JWTMiddleware())
	e.POST("/users", userHandlerAPI.CreateNewUser, middlewares.JWTMiddleware())
}
