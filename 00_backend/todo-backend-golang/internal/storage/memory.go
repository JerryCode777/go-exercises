package storage

import (
	"errors"
	"sync"
	"todo-backend-golang/internal/models"
)

// MemoryStorage	implementa un almacenamiento en memoria para tareas
type MemoryStorage struct {
	tasks  map[int]*models.Task // almacenamietno de tareas
	mu     sync.Mutex           //mutex para proteger el accesos concurrente
	nextID int                  // id para la proxima tarea
}

// create agrega una nueva tarea en el almacenamiento
func (s *MemoryStorage) Create(task *models.Task) *models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task.ID = s.nextID
	s.tasks[s.nextID] = task
	s.nextID++
	return task
}

// GetAll devuelve todas las tareas
func (s *MemoryStorage) GetAll() []*models.Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	var result []*models.Task
	for _, task := range s.tasks {
		result = append(result, task)
	}
	return result
}

// GetByID devuelve una tarea por su ID
func (s *MemoryStorage) GetByID(id int) (*models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

// Update actualiza una tarea existente
func (s *MemoryStorage) Update(id int, updateTask *models.Task) (*models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	task.Title = updateTask.Title
	task.Completed = updateTask.Completed
	return task, nil
}

// Delete elimina una tarea por su ID
func (s *MemoryStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(s.tasks, id)
	return nil
}

