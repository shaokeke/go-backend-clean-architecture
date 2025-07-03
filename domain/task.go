package domain

import (
	"context"
)

type Task struct {
	ID     uint   `json:"-" gorm:"primaryKey;autoIncrement"`
	Title  string `json:"title" binding:"required"  gorm:"type:varchar(128);not null"`
	UserID uint   `json:"-" gorm:"comment:用户ID"`
}

type TaskRepository interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID uint) ([]Task, error)
}

type TaskUsecase interface {
	Create(c context.Context, task *Task) error
	FetchByUserID(c context.Context, userID uint) ([]Task, error)
}
