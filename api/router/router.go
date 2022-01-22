package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"

	"github.com/YutoSekiguchi/todo/controller"
)

func InitRouter(db *gorm.DB) {
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:7120", "http://localhost:8080"},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	ctrl := controller.NewController(db)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Golang!!")
	})

	// Task
	task := e.Group("/tasks")
	{
		task.GET("", ctrl.HandleGetTaskList)
		task.POST("", ctrl.HandlePostTask)
		task.GET("/:id", ctrl.HandleGetTaskById)
		task.PUT("/edit/:id", ctrl.HandleUpdateTask)
		task.DELETE("/delete/:id", ctrl.HandleDeleteTask)
	}

	// Start server
	e.Logger.Fatal(e.Start(":9000"))
}