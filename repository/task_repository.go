package repository

import (
	"context"
	"gorm.io/gorm"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
)

type taskRepository struct {
	database *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepository{
		database: db,
	}
}

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	err := tr.database.Create(&task).Error
	return err
}

func (tr *taskRepository) FetchByUserID(c context.Context, userID uint) ([]domain.Task, error) {
	var tasks []domain.Task
	err := tr.database.Find(&tasks, "user_id = ?", userID).Error
	return tasks, err
}
