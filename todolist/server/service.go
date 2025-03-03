package main

import (
	"context"
	
	"fmt"


	pb "example.com/Todo/todolist/proto"
	"example.com/Todo/ai"
	"example.com/Todo/database"
)

type TodoServiceServer struct {
	pb.UnimplementedTodoServiceServer
}
 
func (s *TodoServiceServer) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	query := "INSERT INTO tasks (title, description, due_date, status) VALUES ($1, $2, $3, 'Pending') RETURNING id"
	var id int
	err := database.DB.QueryRow(query, req.Title, req.Description, req.DueDate).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &pb.CreateTaskResponse{
		Task: &pb.Task{
			Id:          int32(id),
			Title:       req.Title,
			Description: req.Description,
			Status:      "Pending",
			DueDate:     req.DueDate,
		},
	}, nil
}


func (s *TodoServiceServer) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.GetTaskResponse, error) {
	query := "SELECT id, title, description, status, due_date FROM tasks WHERE id = $1"
	var task pb.Task
	err := database.DB.QueryRow(query, req.Id).Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.DueDate)
	if err != nil {
		return nil, err
	}

	return &pb.GetTaskResponse{Task: &task}, nil
}


func (s *TodoServiceServer) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	query := "SELECT id, title, description, status, due_date FROM tasks"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*pb.Task
	for rows.Next() {
		var task pb.Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.DueDate); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return &pb.ListTasksResponse{Tasks: tasks}, nil
}


func (s *TodoServiceServer) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	query := "UPDATE tasks SET title=$1, description=$2, status=$3, due_date=$4 WHERE id=$5"
	_, err := database.DB.Exec(query, req.Title, req.Description, req.Status, req.DueDate, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTaskResponse{
		Task: &pb.Task{
			Id:          req.Id,
			Title:       req.Title,
			Description: req.Description,
			Status:      req.Status,
			DueDate:     req.DueDate,
		},
	}, nil
}

 
func (s *TodoServiceServer) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*pb.DeleteTaskResponse, error) {
	query := "DELETE FROM tasks WHERE id=$1"
	_, err := database.DB.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteTaskResponse{Message: "Task deleted successfully"}, nil
}

// AI-Powered Suggestions
func (s *TodoServiceServer) GetAISuggestions(ctx context.Context, req *pb.AISuggestionsRequest) (*pb.AISuggestionsResponse, error) {
	var tasksStr string
	for _, task := range req.Tasks {
		tasksStr += fmt.Sprintf("Title: %s, Status: %s, Due: %s | ", task.Title, task.Status, task.DueDate)
	}

	suggestion, err := ai.GetAISuggestions(tasksStr)
	if err != nil {
		return nil, err
	}

	return &pb.AISuggestionsResponse{Suggestion: suggestion}, nil
}