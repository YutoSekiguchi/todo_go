package controller

import (
	"fmt"
	"net/http"

	"github.com/YutoSekiguchi/todo/service"
	"github.com/labstack/echo/v4"
)

// HandleGetTaskList GET /tasks タスク一覧の取得
func (ctrl Controller) HandleGetTaskList(c echo.Context) error {
	var s service.TaskService
	p, err := s.GetTaskList(ctrl.Db)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandlePostTask POST /tasks タスクの追加
func (ctrl Controller) HandlePostTask(c echo.Context) error {
	var s service.TaskService
	p, err := s.PostTask(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleGetTaskById GET /tasks/:id 指定したidのタスクの取得
func (ctrl Controller) HandleGetTaskById(c echo.Context) error {
	var s service.TaskService
	p, err := s.GetTaskById(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleUpdateTask UPDATE /tasks/edit/:id 指定したidのタスクを更新
func (ctrl Controller) HandleUpdateTask(c echo.Context) error {
	var s service.TaskService
	p, err := s.UpdateTask(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}

// HandleDeleteTask DELETE /tasks/delete/:id 指定したidのタスクを削除
func (ctrl Controller) HandleDeleteTask(c echo.Context) error {
	var s service.TaskService
	p, err := s.DeleteTask(ctrl.Db, c)

	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, err.Error())
	} else {
		return c.JSON(200, p)
	}
}