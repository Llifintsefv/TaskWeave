package handlers

import (
	"TaskWeave/pkg/models"
	pb "TaskWeave/proto/TaskWeave"
	"TaskWeave/server/service"
	"context"
	"fmt"
	"log"
)

type TaskServer struct {
	pb.UnimplementedTaskServiceServer
	service service.TaskService
}

func NewHandlerServer(service service.TaskService) *TaskServer {
	return &TaskServer{service: service}
}

func (s *TaskServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	task := &models.Task{Name: req.Name, Description: req.Description}
	id, err := s.service.CreateTask(ctx, task)
	if err != nil {
		log.Printf("Error creating task: %v", err)
		return nil, err
	}
	task.ID = uint(id) // Приведение к uint
	grpcTask := toGrpcTask(task) //Новая функция
	return &pb.CreateTaskResponse{Task: grpcTask}, nil
}

func (s *TaskServer) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	task, err := s.service.GetTask(ctx, int(req.Id))
	if err != nil {
		log.Printf("Error getting task: %v", err)
		return nil, err
	}
	grpcTask := toGrpcTask(task) //Новая функция
	return &pb.GetTaskResponse{Task: grpcTask}, nil

}

func (s *TaskServer) GetAllTasks(ctx context.Context, req *pb.GetAllTasksRequest) (*pb.GetAllTasksResponse, error) {
	tasks, err := s.service.GetAllTasks(ctx)
	if err != nil {
		log.Printf("Error getting all tasks: %v", err)
		return nil, err
	}
	grpcTasks := make([]*pb.Task, len(tasks))
	for i, task := range tasks {
		grpcTasks[i] = toGrpcTask(task)
	}
	return &pb.GetAllTasksResponse{Tasks: grpcTasks}, nil
}

func (s *TaskServer) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	err := s.service.DeleteTask(ctx, int(req.Id))
	if err != nil {
		log.Printf("Error deleting task: %v", err)
		return nil, err
	}
	return &pb.DeleteTaskResponse{}, nil
}

func (s *TaskServer) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	task := &models.Task{
		Name:        req.Name,
		Description: req.Description,
	}

	existingTask, err := s.service.GetTask(ctx, int(req.Id))
	if err != nil {
		log.Printf("Error getting task for update: %v", err)
		return nil, err
	}
	if existingTask == nil {
		return nil, fmt.Errorf("task with ID %d not found", req.Id)
	}
	existingTask.Name = task.Name
	existingTask.Description = task.Description

	err = s.service.UpdateTask(ctx, existingTask) 
	if err != nil {
		log.Printf("Error updating task: %v", err)
		return nil, err
	}

	grpcTask := toGrpcTask(existingTask)
	return &pb.UpdateTaskResponse{Task: grpcTask}, nil
}

func toGrpcTask(task *models.Task) *pb.Task {
	if task == nil {
		return nil
	}
	return &pb.Task{Id: int32(task.ID), Name: task.Name, Description: task.Description}
}