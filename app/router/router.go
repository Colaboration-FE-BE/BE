package router

import (
	"net/http"

	"immersive-dash-4/app/middlewares"
	_userData "immersive-dash-4/features/user/data"
	_userHandler "immersive-dash-4/features/user/handler"
	_userService "immersive-dash-4/features/user/service"

	_classData "immersive-dash-4/features/class/data"
	_classHandler "immersive-dash-4/features/class/handler"
	_classService "immersive-dash-4/features/class/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := _userData.New(db)
	userService := _userService.New(userData)
	userHandlerAPI := _userHandler.New(userService)

	classData := _classData.New(db)
	classService := _classService.New(classData)
	classHandlerAPI := _classHandler.New(classService)

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
	e.GET("/users", userHandlerAPI.GetAllUsers, middlewares.JWTMiddleware())
	e.DELETE("/users/:user_id", userHandlerAPI.DeleteUser, middlewares.JWTMiddleware())
	e.PUT("/users/:user_id", userHandlerAPI.UpdateUser, middlewares.JWTMiddleware())

	//Class
	e.GET("/class", classHandlerAPI.GetAllClass, middlewares.JWTMiddleware())
	e.POST("/class", classHandlerAPI.CreateClass, middlewares.JWTMiddleware())
	e.GET("/class/:class_id", classHandlerAPI.GetClassById, middlewares.JWTMiddleware())
	e.DELETE("/class/:class_id", classHandlerAPI.DeleteClass, middlewares.JWTMiddleware())
}
