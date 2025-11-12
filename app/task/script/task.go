package script

import (
	"context"
	"go-micro-todoList/app/task/repository/mq/task"
)

func TaskCreateSync(ctx context.Context) {
	tSync := new(task.SyncTask)
	err := tSync.RunTaskService(ctx)
	if err != nil {
		return
	}
}
