//cria um struct que representa todas as rotas da API

package router

import (
	"api/src/controllers"

	"github.com/labstack/echo/v4"
)

func UserRoute(e *echo.Echo)  {
    e.POST("/user", controllers.CreateUser)
    e.GET("/user/:userId", controllers.GetAUser)
    e.PUT("/user/:userId", controllers.EditAUser)
    e.DELETE("/user/:userId", controllers.DeleteAUser)
    e.GET("/users", controllers.GetAllUsers)
}