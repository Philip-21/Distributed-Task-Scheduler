package scheduler

import "sync"

//manages tasks to be scheduled

type TaskManager struct {
	Wg sync.WaitGroup
}

func NewTaskManager() *TaskManager {
	t := &TaskManager{}
	return t
}

func (t*TaskManager) ScheduleOneTimeTask(taskId string){}