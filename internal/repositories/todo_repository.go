package repositories

import (
	"errors"

	"todo-api/internal/models"

	"gorm.io/gorm"
)

var (
	ErrNotFound = errors.New("recurso no encontrado")
	ErrConflict = errors.New("conflicto: el recurso ya existe")
)

type TodoRepository struct {
	DB *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

// Obtiene todos los To Do's existentes
func (r *TodoRepository) GetAll() ([]models.Todo, error) {
	var todos []models.Todo
	if err := r.DB.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}

// Crea un nuevo To Do
func (r *TodoRepository) Create(todo *models.Todo) error {
	// Validar que el titulo no este vacio
	if todo.Title == "" {
		return errors.New("el titulo es requerido")
	}

	// Validar que el To Do no exista ya
	existingTodo, err := r.GetByTitle(todo.Title)
	if err != nil && !errors.Is(err, ErrNotFound) {
		return err
	}
	if existingTodo != nil {
		return ErrConflict
	}

	// Crear el todo
	return r.DB.Create(todo).Error
}

// Obtiene un To Do por su ID
func (r *TodoRepository) GetByID(id string) (*models.Todo, error) {
	var todo models.Todo
	if err := r.DB.First(&todo, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &todo, nil
}

// Actualiza un To Do existente
func (r *TodoRepository) Update(id string, updates map[string]interface{}) error {
	result := r.DB.Model(&models.Todo{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

// Elimina un To Do Existente
func (r *TodoRepository) Delete(id string) error {
	result := r.DB.Delete(&models.Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

// Obtiene un To Do por su titulo
func (r *TodoRepository) GetByTitle(title string) (*models.Todo, error) {
	var todo models.Todo
	if err := r.DB.Where("title = ?", title).First(&todo).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNotFound
		}
		return nil, err
	}
	return &todo, nil
}
