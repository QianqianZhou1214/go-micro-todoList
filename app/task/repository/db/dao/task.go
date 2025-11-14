package dao

import (
	"context"
	"go-micro-todoList/app/task/repository/db/model"
	"go-micro-todoList/idl/pb"

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

func (dao *TaskDao) ListTaskByUserId(userId uint64, start, limit int) (r []*model.Task, count int64, err error) {
	err = dao.Model(&model.Task{}).Offset(start).Limit(limit).Where("uid=?", userId).Find(&r).Error
	if err != nil {
		return
	}

	err = dao.Model(&model.Task{}).Where("uid=?", userId).Count(&count).Error
	return
}

func (dao *TaskDao) GetTaskByIdAndUserId(tId, uId uint64) (r *model.Task, err error) {
	err = dao.Model(&model.Task{}).Where("id=? AND uid = ?", tId, uId).Find(&r).Error
	return
}

func (dao *TaskDao) UpdateTask(req *pb.TaskRequest) (err error) {
	var r *model.Task
	err = dao.Model(&model.Task{}).Where("id=? AND uid = ?", req.Id, req.Uid).Find(&r).Error
	if err != nil {
		return
	}
	r.Title = req.Title
	r.Status = int(req.Status)
	r.Content = req.Content

	return dao.Save(&r).Error

}

func (dao *TaskDao) DeleteTaskByIdAndUserId(tId, uId uint64) (err error) {
	err = dao.Model(&model.Task{}).Where("id=? AND uid = ?", tId, uId).Delete(&model.Task{}).Error
	return
}
