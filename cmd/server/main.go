package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-app/internal/task"
	"todo-app/internal/storage"
)

func main() {
	// Initialize task service (in-memory store or database)
	store := storage.NewInMemoryStore()
	service := task.NewService(store)

	// Set up HTTP handlers
	http.HandleFunc("/tasks", task.GetTasksHandler(service))           // GET all tasks
	http.HandleFunc("/tasks/", task.GetTaskHandler(service))           // GET specific task by ID
	http.HandleFunc("/tasks/create", task.CreateTaskHandler(service))  // POST to create a task
	http.HandleFunc("/tasks/update/", task.UpdateTaskHandler(service)) // PUT to update a task
	http.HandleFunc("/tasks/delete/", task.DeleteTaskHandler(service)) // DELETE a task

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
