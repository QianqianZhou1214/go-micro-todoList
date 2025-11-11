package dao

import (
	"context"
	"go-micro-todoList/app/task/repository/db/model"

	"gorm.io/gorm"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(ctx context.Context) *TaskDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &TaskDao{NewDBClient(ctx)}
}

func (dao *TaskDao) CreateTask(data *model.Task) error {
	return dao.Model(&model.Task{}).Create(data).Error
}
