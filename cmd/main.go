package main

import (
	"TaskWeave/pkg/db"
	"TaskWeave/server/handlers"
	"TaskWeave/server/repository"
	"TaskWeave/server/service"
	"log"
	"net"
	pb "TaskWeave/proto/TaskWeave"
	"google.golang.org/grpc"
)

func main() {
	
	db,err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	taskRepo := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepo)
	taskHandler := handlers.NewHandlerServer(taskService)

	lis, err := net.Listen("tcp", ":50051")
	grpcServer := grpc.NewServer()
	pb.RegisterTaskServiceServer(grpcServer, taskHandler)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}