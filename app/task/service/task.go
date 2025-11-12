package service

import (
	"context"
	"encoding/json"
	"go-micro-todoList/app/task/repository/db/dao"
	"go-micro-todoList/app/task/repository/db/model"
	"go-micro-todoList/app/task/repository/mq"
	"go-micro-todoList/idl/pb"
	"go-micro-todoList/pkg/e"
	"sync"
)

type TaskServ struct {
}

var TaskSevIns *TaskServ
var TaskSrvOnce sync.Once

// GetTaskServ Lazy Singleton Pattern
func GetTaskServ() *TaskServ {
	TaskSrvOnce.Do(func() {
		TaskSevIns = &TaskServ{}
	})
	return TaskSevIns
}

// create task sending to MQ
func (t *TaskServ) CreateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	body, _ := json.Marshal(req)
	err = mq.SendMessage2MQ(body)
	if err != nil {
		resp.Code = e.Error
		return
	}
	return
}

func TaskMQ2DB(ctx context.Context, req *pb.TaskRequest) (err error) {
	m := &model.Task{
		Uid:       uint(req.Uid),
		Title:     req.Title,
		Status:    int(req.Status),
		Content:   req.Content,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
	}
	return dao.NewTaskDao(ctx).CreateTask(m)

}
