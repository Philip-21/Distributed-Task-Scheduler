// manages tasks to be scheduled
package scheduler

import (
	"log"
	"sync"
	"time"

	"github.com/Philip-21/dts/internal/common"
	"github.com/Philip-21/dts/internal/worker"
	"google.golang.org/protobuf/types/known/anypb"
)

type TaskManager struct {
	Wg     sync.WaitGroup
	Mt     sync.Mutex
	Tasks  []common.Task
	Worker *worker.Worker
}

func NewTaskManager(worker *worker.Worker) *TaskManager {
	t := &TaskManager{
		Worker: worker,
		Tasks:  make([]common.Task, 0),
	}
	return t
}

func (t *TaskManager) OneTimeTask(taskId string, payload *anypb.Any, executeAt string) (*common.Task, error) {
	execTime, err := time.Parse(time.RFC3339, executeAt)
	if err != nil {
		return nil, err
	}
	task := &common.Task{
		ID:        taskId,
		Payload:   payload,
		ExecuteAt: execTime,
	}
	t.Mt.Lock()
	t.Tasks = append(t.Tasks, *task)
	t.Mt.Unlock()

	// Schedule the task execution
	go func() {
		time.Sleep(time.Until(task.ExecuteAt))
		executedTask, err := t.Worker.ExecuteOneTimeTask(task.ID)
		if err != nil {
			log.Printf("error executing task: %v", err)
			return
		}
		log.Printf("task executed successfully: %+v", executedTask)
	}()

	return task, nil
}

func (t *TaskManager) ScheduleRecuringTask() {}
