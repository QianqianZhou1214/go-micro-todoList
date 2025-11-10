package service

import (
	"context"
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

func (t *TaskServ) CreateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	return
}
