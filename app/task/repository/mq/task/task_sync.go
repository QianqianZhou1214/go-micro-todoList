package task

import (
	"context"
	"encoding/json"
	"go-micro-todoList/app/task/repository/mq"
	"go-micro-todoList/app/task/service"
	"go-micro-todoList/consts"
	"go-micro-todoList/idl/pb"
)

type SyncTask struct {
}

func (s *SyncTask) RunTaskService(ctx context.Context) (err error) {
	rabbitMqQueue := consts.RabbitMqTaskQueue
	msgs, err := mq.ConsumeMessage(ctx, rabbitMqQueue)
	if err != nil {
		return
	}
	var forever chan struct{} // for blocking
	go func() {
		for d := range msgs {
			req := new(pb.TaskRequest)
			err = json.Unmarshal(d.Body, req)
			if err != nil {
				return
			}
			err = service.TaskMQ2DB(ctx, req)
			if err != nil {
				return
			}
		}
	}()
	<-forever
	return nil
}
