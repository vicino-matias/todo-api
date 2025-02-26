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
	if todo.Title == "" {
		return errors.New("el título es requerido")
	}

	// Verifica si el To Do ya existe
	exists, err := r.ExistsByTitle(todo.Title)
	if err != nil {
		return err
	}
	if exists {
		return ErrConflict
	}

	return r.DB.Create(todo).Error
}

// Obtiene un To Do por su ID
func (r *TodoRepository) GetByID(id string) (*models.Todo, error) {
	var todo models.Todo
	if err := r.DB.First(&todo, "id = ?", id).Error; err != nil {
		return nil, r.handleDBError(err)
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

// Elimina un To Do existente
func (r *TodoRepository) Delete(id string) error {
	result := r.DB.Delete(&models.Todo{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

// Obtiene un To Do por su título
func (r *TodoRepository) GetByTitle(title string) (*models.Todo, error) {
	var todo models.Todo
	if err := r.DB.Where("title = ?", title).First(&todo).Error; err != nil {
		return nil, r.handleDBError(err)
	}
	return &todo, nil
}

// Verifica si un To Do con el mismo título ya existe
func (r *TodoRepository) ExistsByTitle(title string) (bool, error) {
	var count int64
	if err := r.DB.Model(&models.Todo{}).Where("title = ?", title).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

// Maneja errores de base de datos y los traduce a errores del repositorio
func (r *TodoRepository) handleDBError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	return err
}
