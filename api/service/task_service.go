package service

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type TaskService struct{}

// タスク一覧を取得
func (s TaskService) GetTaskList(db *gorm.DB) ([]Task, error) {
	var task []Task

	if err := db.Find(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

// PostTask タスクの追加
// request bodyでjsonを渡す
func (s TaskService) PostTask(db *gorm.DB, c echo.Context) (Task, error) {
	var task Task
	c.Bind(&task)

	if err := db.Create(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

// idを指定してそのタスクの抽出
func (s TaskService) GetTaskById(db *gorm.DB, c echo.Context) ([]Task, error) {
	var task []Task
	id := c.Param("id")
	
	if err := db.First(&task, id).Scan(&task).Error; err != nil {
		return task, err
	}
	return task, nil
}

// idを指定してそのタスクの更新
func (s TaskService) UpdateTask(db *gorm.DB, c echo.Context) (Task, error) {
	var task Task
	id := c.Param("id")
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		return task, err
	}
	if err := c.Bind(&task); err != nil {
		return task, err
	}
	db.Save(&task)

	return task, nil
}

// idを指定してそのタスクの削除
func (s TaskService) DeleteTask(db *gorm.DB, c echo.Context) ([]Task, error) {
	var task []Task
	id := c.Param("id")

	if err := db.Where("id = ?", id).Delete(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}