package worker

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Philip-21/dts/internal/common"
	"google.golang.org/protobuf/types/known/anypb"
)

type Worker struct {
	Mt    sync.Mutex
	Tasks []common.Task
}

type Task struct {
	ID        string
	Payload   *anypb.Any
	ExecuteAt time.Time
}

func NewTaskWorker() *Worker {
	w := &Worker{
		Tasks: make([]common.Task, 0),
	}
	return w
}

func (t *Worker) FetchTaskByID(taskID string) (*common.Task, error) {
	t.Mt.Lock()
	defer t.Mt.Unlock()
	for _, task := range t.Tasks {
		if task.ID == taskID {
			return &task, nil
		}
	}
	return nil, fmt.Errorf("task with ID %s not found", taskID)
}

func (w *Worker) ExecuteOneTimeTask(taskID string) (*common.Task, error) {

	task, err := w.FetchTaskByID(taskID)
	if err != nil {
		return nil, err
	}

	// Implement the task execution logic
	log.Printf("executing task: %+v", task)

	// Simulate task execution
	time.Sleep(2 * time.Second)

	// Remove the task from the queue after execution
	for i, tsk := range w.Tasks {
		if tsk.ID == taskID {
			w.Tasks = append(w.Tasks[:i], w.Tasks[i+1:]...)
			break
		}

	}

	fmt.Printf("Executing one-time task %s with payload: %v\n", task.ID, task.Payload)
	return task, nil
}
