package task

import (
	"errors"
	"todo-app/internal/storage"
)

// Service handles business logic related to tasks
type Service struct {
	store storage.Store
}

// NewService creates a new Task service
func NewService(store storage.Store) *Service {
	return &Service{
		store: store,
	}
}

// CreateTask creates a new task
func (s *Service) CreateTask(task Task) (Task, error) {
	return s.store.Create(task)
}

// GetTasks returns all tasks
func (s *Service) GetTasks() ([]Task, error) {
	return s.store.GetAll()
}

// GetTaskByID retrieves a task by its ID
func (s *Service) GetTaskByID(id int) (Task, error) {
	return s.store.GetByID(id)
}

// UpdateTask updates an existing task
func (s *Service) UpdateTask(id int, task Task) (Task, error) {
	return s.store.Update(id, task)
}

// DeleteTask deletes a task
func (s *Service) DeleteTask(id int) error {
	return s.store.Delete(id)
}
