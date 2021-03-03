package Server

import (
	pb "./proto/"
	"context"
	"github.com/google/uuid"
	"log"
)

type server struct {
	tasks []*pb.TaskObj
}

func (s *server) CreateTask(ctx context.Context, newTask *pb.CreateRequest) (*pb.TaskObj, error) {
	log.Printf("Received new task %s", newTask)
	taskObj := &pb.TaskObj{
		Id:          uuid.New().String(),
		Description: newTask.Description,
	}
	s.tasks = append(s.tasks, taskObj)
	return taskObj, nil
}

func (s *server) RemoveTask(ctx context.Context, taskToRemove *pb.RemoveRequest) (*pb.Response, error) {
	message := "Unsuccessful"
	for i, task := range s.tasks {
		if task.Id == taskToRemove.Id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			message = "Successful"
			break
		}
	}
	return &pb.Response{Message: message}, nil
}
func (s *server) UpdateTask(ctx context.Context, taskToUpdate *pb.TaskObj) (*pb.Response, error) {
	message := "Unsuccessful"
	for _, task := range s.tasks {
		if task.Id == taskToUpdate.Id {
			task.Description = taskToUpdate.Description
			message = "Successful"
			break
		}
	}
	return &pb.Response{Message: message}, nil
}
