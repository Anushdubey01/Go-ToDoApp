package storage

import (
	"errors"
	"todo-app/internal/task"
)

// InMemoryStore is an implementation of the Store interface using an in-memory map
type InMemoryStore struct {
	tasks map[int]task.Task
}

// NewInMemoryStore creates a new in-memory task store
func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		tasks: make(map[int]task.Task),
	}
}

// Create adds a new task to the store
func (s *InMemoryStore) Create(task task.Task) (task.Task, error) {
	task.ID = len(s.tasks) + 1
	s.tasks[task.ID] = task
	return task, nil
}

// GetAll returns all tasks
func (s *InMemoryStore) GetAll() ([]task.Task, error) {
	var tasks []task.Task
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

// GetByID retrieves a task by its ID
func (s *InMemoryStore) GetByID(id int) (task.Task, error) {
	if task, exists := s.tasks[id]; exists {
		return task, nil
	}
	return task.Task{}, errors.New("task not found")
}

// Update modifies an existing task
func (s *InMemoryStore) Update(id int, updatedTask task.Task) (task.Task, error) {
	if _, exists := s.tasks[id]; !exists {
		return task.Task{}, errors.New("task not found")
	}
	updatedTask.ID = id
	s.tasks[id] = updatedTask
	return updatedTask, nil
}

// Delete removes a task from the store
func (s *InMemoryStore) Delete(id int) error {
	if _, exists := s.tasks[id]; !exists {
		return errors.New("task not found")
	}
	delete(s.tasks, id)
	return nil
}
