package main

import (
	"context"
	"log"

	pb "TaskWeave/proto/TaskWeave"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


type TaskClient struct {
	client pb.TaskServiceClient
}


func NewTaskClient(address string) (*TaskClient, error) {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewTaskServiceClient(conn)
	return &TaskClient{client: client}, nil
}


func (tc *TaskClient) CreateTask(ctx context.Context, name, description string) (*pb.Task, error) {
	req := &pb.CreateTaskRequest{Name: name, Description: description}
	res, err := tc.client.CreateTask(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Task, nil
}

func (tc *TaskClient) GetTask(ctx context.Context, id int) (*pb.Task, error) {
	req := &pb.GetTaskRequest{Id: int32(id)}
	res, err := tc.client.GetTask(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Task, nil
}


func (tc *TaskClient) GetAllTasks(ctx context.Context) ([]*pb.Task, error) {
	req := &pb.GetAllTasksRequest{}
	res, err := tc.client.GetAllTasks(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Tasks, nil
}


func (tc *TaskClient) UpdateTask(ctx context.Context, id int, name, description string) (*pb.Task, error) {
	req := &pb.UpdateTaskRequest{Id: int32(id), Name: name, Description: description}
	res, err := tc.client.UpdateTask(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Task, nil
}


func (tc *TaskClient) DeleteTask(ctx context.Context, id int) error {
	req := &pb.DeleteTaskRequest{Id: int32(id)}
	_, err := tc.client.DeleteTask(ctx, req)
	return err
}

func main() {
	taskClient, err := NewTaskClient("localhost:50051")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	ctx := context.Background()

	task, err := taskClient.CreateTask(ctx, "test task", "Description Task")
	if err != nil {
		log.Printf("error creating: %v", err)
	} else {
		log.Printf("created task: %+v", task)
	}

	
	task, err = taskClient.GetTask(ctx, int(task.Id)) 
	if err != nil {
		log.Printf("Error get task %d: %v", task.Id, err)
	} else {
		log.Printf("get task: %+v", task)
	}

	tasks, err := taskClient.GetAllTasks(ctx)
	if err != nil {
		log.Printf("Error get all tasks: %v", err)
	} else {
		log.Printf("all Tasks:")
		for _, task := range tasks {
			log.Printf(" - %+v", task)
		}
	}


	updatedTask, err := taskClient.UpdateTask(ctx, int(task.Id), "Updated Ttsk", "Updated description")
	if err != nil {
		log.Printf("Error updating task %d: %v", task.Id, err)
	} else {
		log.Printf("Updated Task: %+v", updatedTask)
	}


	err = taskClient.DeleteTask(ctx, int(task.Id))
	if err != nil {
		log.Printf("Error delet task %d: %v", task.Id, err)
	} else {
		log.Printf("Deleted Task %d", task.Id)
	}


    task, err = taskClient.GetTask(ctx, int(task.Id))
    if err != nil {
        log.Printf("Error get task %d: %v", task.Id, err)
    } else {
        log.Printf("Get Task: %+v", task)
    }


	tasks, err = taskClient.GetAllTasks(ctx)
	if err != nil {
		log.Printf("Error get all tasks: %v", err)
	} else {
		log.Printf("All tasks:")
		for _, task := range tasks {
			log.Printf(" - %+v", task)
		}
	}
}