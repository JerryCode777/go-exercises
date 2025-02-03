// internal/storage/memory.go
package storage

import (
	"errors"
	"sync"
	"todo-backend-golang/internal/models"
)

// MemoryStorage implementa un almacenamiento en memoria para tareas
type MemoryStorage struct {
	Tasks  map[int]*models.Task // Cambiado a mayúscula para exportarlo
	Mu     sync.Mutex           // Cambiado a mayúscula si se necesita usar desde otro paquete
	NextID int                  // Cambiado a mayúscula para exportarlo
}

// NewMemoryStorage crea una nueva instancia de almacenamiento en memoria
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		Tasks:  make(map[int]*models.Task),
		NextID: 1,
	}
}

// Create agrega una nueva tarea en el almacenamiento
func (s *MemoryStorage) Create(task *models.Task) *models.Task {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	task.ID = s.NextID
	s.Tasks[s.NextID] = task
	s.NextID++
	return task
}

// GetAll devuelve todas las tareas
func (s *MemoryStorage) GetAll() []*models.Task {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	var result []*models.Task
	for _, task := range s.Tasks {
		result = append(result, task)
	}
	return result
}

// GetByID devuelve una tarea por su ID
func (s *MemoryStorage) GetByID(id int) (*models.Task, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	task, exists := s.Tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return task, nil
}

// Update actualiza una tarea existente
func (s *MemoryStorage) Update(id int, updateTask *models.Task) (*models.Task, error) {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	task, exists := s.Tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	task.Title = updateTask.Title
	task.Completed = updateTask.Completed
	return task, nil
}

// Delete elimina una tarea por su ID
func (s *MemoryStorage) Delete(id int) error {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	_, exists := s.Tasks[id]
	if !exists {
		return errors.New("task not found")
	}
	delete(s.Tasks, id)
	return nil
}
