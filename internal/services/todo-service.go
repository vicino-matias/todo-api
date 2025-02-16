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
	todoRepo repositories.TodoRepository
}

func NewTodoService(todoRepo repositories.TodoRepository) *TodoService {
	return &TodoService{todoRepo: todoRepo}
}

// Obtiene todo los los To Do's existentes
func (s *TodoService) GetAllTodos() ([]models.Todo, error) {
	return s.todoRepo.GetAll()
}

// Crea un nuevo To Do

func (s *TodoService) CreateTodo(todo *models.Todo) error {
	// Valida que el titulo no este vacio
	if todo.Title == "" {
		return errors.New("el titulo es requerido")
	}

	// Validar que el to do no exista ya
	existingTodo, err := s.todoRepo.GetByTitle(todo.Title)
	if err != nil && !errors.Is(err, repositories.ErrNotFound) {
		return err
	}
	if existingTodo != nil {
		return ErrConflict
	}

	// Crear el To Do
	return s.todoRepo.Create(todo)
}

// Obtiene un To Do por su ID
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

// Actualiza un todo existente
func (s *TodoService) UpdateTodo(id string, updates map[string]interface{}) error {
	// Verifica si el to do existe
	if _, err := s.todoRepo.GetByID(id); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return ErrNotFound
		}
		return err
	}
	// Actualiza el to do
	return s.todoRepo.Update(id, updates)
}

// Elimina un To Do por su ID
func (s *TodoService) DeleteTodo(id string) error {
	// Verificar si el todo existe
	if _, err := s.todoRepo.GetByID(id); err != nil {
		if errors.Is(err, repositories.ErrNotFound) {
			return ErrNotFound
		}
		return err
	}

	// Eliminar el to do
	return s.todoRepo.Delete(id)
}
