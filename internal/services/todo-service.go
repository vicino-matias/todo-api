package services

import (
	"errors"
	"todo-api/internal/models"
	"todo-api/internal/repositories"
)

var (
	ErrNotFound = errors.New("recurso no encontrado")
	ErrConflict = errors.New("conflicto: el recurso ya existe")
)

type TodoService struct {
	todoRepo *repositories.TodoRepository
}

// NewTodoService crea una nueva instancia del servicio de To Do's
func NewTodoService(todoRepo *repositories.TodoRepository) *TodoService {
	return &TodoService{todoRepo: todoRepo}
}

// GetAllTodos obtiene todos los To Do's existentes
func (s *TodoService) GetAllTodos() ([]models.Todo, error) {
	return s.todoRepo.GetAll()
}

// CreateTodo valida y crea un nuevo To Do
func (s *TodoService) CreateTodo(todo *models.Todo) error {
	if todo.Title == "" {
		return errors.New("el título es requerido")
	}

	// Verifica si ya existe un To Do con el mismo título
	existingTodo, err := s.todoRepo.GetByTitle(todo.Title)
	if err != nil && !errors.Is(err, repositories.ErrNotFound) {
		return err
	}
	if existingTodo != nil {
		return ErrConflict
	}

	return s.todoRepo.Create(todo)
}

// GetTodoByID obtiene un To Do por su ID
func (s *TodoService) GetTodoByID(id string) (*models.Todo, error) {
	todo, err := s.todoRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return todo, nil
}

// UpdateTodo actualiza un To Do existente
func (s *TodoService) UpdateTodo(id string, updates map[string]interface{}) error {
	if err := s.validateExistenceByID(id); err != nil {
		return err
	}
	return s.todoRepo.Update(id, updates)
}

// DeleteTodo elimina un To Do por su ID
func (s *TodoService) DeleteTodo(id string) error {
	if err := s.validateExistenceByID(id); err != nil {
		return err
	}
	return s.todoRepo.Delete(id)
}

// validateExistenceByID verifica si un To Do existe antes de operar sobre él
func (s *TodoService) validateExistenceByID(id string) error {
	_, err := s.todoRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return ErrNotFound
		}
		return err
	}
	return nil
}
