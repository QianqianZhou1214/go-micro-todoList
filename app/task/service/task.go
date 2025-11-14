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

var TaskSevIns *TaskServ
var TaskSrvOnce sync.Once

type TaskServ struct {
}

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

func (*TaskServ) GetTasksList(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskListResponse) (err error) {
	resp.Code = e.Success
	if req.Limit == 0 {
		req.Limit = 10
	}
	r, count, err := dao.NewTaskDao(ctx).ListTaskByUserId(req.Uid, int(req.Start), int(req.Limit))
	if err != nil {
		resp.Code = e.Error
		return
	}
	var taskRes []*pb.TaskModel
	for _, item := range r {
		taskRes = append(taskRes, BuildTask(item))
	}
	resp.TaskList = taskRes
	resp.Count = uint32(count)
	return
}

func (*TaskServ) GetTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	r, err := dao.NewTaskDao(ctx).GetTaskByIdAndUserId(req.Id, req.Uid)
	if r.ID == 0 || err != nil {
		resp.Code = e.Error
		return
	}
	resp.TaskDetail = BuildTask(r)
	return
}

func (*TaskServ) UpdateTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	err = dao.NewTaskDao(ctx).UpdateTask(req)
	if err != nil {
		resp.Code = e.Error
		return
	}
	return
}

func (*TaskServ) DeleteTask(ctx context.Context, req *pb.TaskRequest, resp *pb.TaskDetailResponse) (err error) {
	resp.Code = e.Success
	err = dao.NewTaskDao(ctx).DeleteTaskByIdAndUserId(req.Id, req.Uid)
	if err != nil {
		resp.Code = e.Error
		return
	}

	return
}

func BuildTask(item *model.Task) *pb.TaskModel {
	return &pb.TaskModel{
		Id:         uint64(item.ID),
		Uid:        uint64(item.Uid),
		Title:      item.Title,
		Content:    item.Content,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		Status:     int64(item.Status),
		CreateTime: item.CreatedAt.Unix(),
		UpdateTime: item.UpdatedAt.Unix(),
	}
}
