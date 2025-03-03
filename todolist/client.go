package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	pb "example.com/Todo/todolist/proto"

	"google.golang.org/grpc"
)

func main() {
	
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewTodoServiceClient(conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n========= To-Do List CLI =========")
		fmt.Println("1. Create Task")
		fmt.Println("2. Get Task")
		fmt.Println("3. List Tasks")
		fmt.Println("4. Update Task")
		fmt.Println("5. Get AI Suggestions")
		fmt.Println("6. Delete Task")
		fmt.Println("7. Exit")
		fmt.Print("Enter your choice: ")

		choiceStr, _ := reader.ReadString('\n')
		choiceStr = strings.TrimSpace(choiceStr)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Enter Task Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Enter Description: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)

			fmt.Print("Enter Due Date (YYYY-MM-DD): ")
			dueDate, _ := reader.ReadString('\n')
			dueDate = strings.TrimSpace(dueDate)

			createTask(client, title, description, dueDate)

		case 2:
			fmt.Print("Enter Task ID: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)

			getTask(client, int32(id))

		case 3:
			listTasks(client)

		case 4:
			fmt.Print("Enter Task ID: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)

			fmt.Print("Enter Updated Title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Enter Updated Description: ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)

			fmt.Print("Enter Updated Status (Pending/Completed): ")
			status, _ := reader.ReadString('\n')
			status = strings.TrimSpace(status)

			fmt.Print("Enter Updated Due Date (YYYY-MM-DD): ")
			dueDate, _ := reader.ReadString('\n')
			dueDate = strings.TrimSpace(dueDate)

			updateTask(client, int32(id), title, description, status, dueDate)

		case 5:
			getAISuggestions(client)

		case 6:
			fmt.Print("Enter Task ID to Delete: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			id, _ := strconv.Atoi(idStr)

			deleteTask(client, int32(id))

		case 7:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice, please enter a number between 1 and 7.")
		}
	}
}

func createTask(client pb.TodoServiceClient, title, description, dueDate string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.CreateTaskRequest{
		Title:       title,
		Description: description,
		DueDate:     dueDate,
	}

	res, err := client.CreateTask(ctx, req)
	if err != nil {
		log.Fatalf("Error creating task: %v", err)
	}

	fmt.Printf("✅ Task created: %v\n", res.Task)
}

func getTask(client pb.TodoServiceClient, id int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetTaskRequest{Id: id}

	res, err := client.GetTask(ctx, req)
	if err != nil {
		log.Fatalf("Error getting task: %v", err)
	}

	fmt.Printf(" Task retrieved: %v\n", res.Task)
}

func listTasks(client pb.TodoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.ListTasksRequest{}

	res, err := client.ListTasks(ctx, req)
	if err != nil {
		log.Fatalf("Error listing tasks: %v", err)
	}

	fmt.Println("\n All Tasks:")
	for _, task := range res.Tasks {
		fmt.Printf("- %v\n", task)
	}
}

func updateTask(client pb.TodoServiceClient, id int32, title, description, status, dueDate string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.UpdateTaskRequest{
		Id:          id,
		Title:       title,
		Description: description,
		Status:      status,
		DueDate:     dueDate,
	}

	res, err := client.UpdateTask(ctx, req)
	if err != nil {
		log.Fatalf("Error updating task: %v", err)
	}

	fmt.Printf(" Task updated: %v\n", res.Task)
}

func getAISuggestions(client pb.TodoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.AISuggestionsRequest{}

	res, err := client.GetAISuggestions(ctx, req)
	if err != nil {
		log.Fatalf("Error getting AI suggestions: %v", err)
	}

	fmt.Printf("🤖 AI Suggestion: %v\n", res.Suggestion)
}

func deleteTask(client pb.TodoServiceClient, id int32) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.DeleteTaskRequest{Id: id}

	res, err := client.DeleteTask(ctx, req)
	if err != nil {
		log.Fatalf("Error deleting task: %v", err)
	}

	fmt.Printf("Task deleted: %v\n", res.Message)
}