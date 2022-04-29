package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"example-user-crud/internal/controller"
)

func main() {
	container := newContainer()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/alive", func(context echo.Context) error {
		return context.JSON(http.StatusOK, map[string]interface{}{"status": "alive"})
	})

	//Users CRUD
	e.GET("/users/:userId", controller.GetUser(container.UsersManager))
	e.GET("/users", controller.GetUsers(container.UsersManager))
	e.POST("/users", controller.CreateUser(container.UsersManager))
	e.PUT("/users/:userId", controller.ReplaceUser(container.UsersManager))
	e.DELETE("/users/:userId", controller.DeleteUser(container.UsersManager))

	e.POST("/users/:userId/avatar", controller.ReceiveAvatar(container.UsersManager))

	e.Logger.Fatal(e.Start(":1234"))
}
