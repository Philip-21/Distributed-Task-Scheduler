package scheduler

import (
	"context"
	"log"

	pb "github.com/Philip-21/dts/internal/proto"
	"github.com/Philip-21/dts/internal/worker"
)

//This file will define the gRPC server and implement the methods for scheduling tasks

type Server struct {
	pb.UnimplementedSchedulerServer
	TaskManager *TaskManager
}

func NewServer(worker *worker.Worker) *Server {
	return &Server{
		TaskManager: NewTaskManager(worker),
	}
}

func (s *Server) ScheduleOneTimeTask(ctx context.Context, req *pb.ScheduleRequest) (*pb.ScheduleResponse, error) {
	task := req.GetTask()
	log.Printf("received task: %+v\n", task)
	execTask, err := s.TaskManager.OneTimeTask(task.TaskId, task.Payload, task.ExecuteAt)
	if err != nil {
		return &pb.ScheduleResponse{
			Status:  "error",
			Message: err.Error(),
		}, err
	}
	response := &pb.ScheduleResponse{
		Status:  "successful",
		Message: "task scheduled successfully",

		Task: &pb.OneTimeTask{
			TaskId:  execTask.ID,
			Payload: execTask.Payload,
		},
	}
	return response, nil
}
