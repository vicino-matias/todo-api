package handlers

import (
	"errors"
	"net/http"
	"todo-api/internal/models"
	"todo-api/internal/services"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoService services.TodoService
}

func NewTodoHandler(todoService services.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

// GetTodos obtiene todos los To Dos
func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.todoService.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los To Do's"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// CreateTodo crea un nuevo To Do
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	if todo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El campo 'title' es requerido"})
		return
	}
	if err := h.todoService.CreateTodo(&todo); err != nil {
		status := http.StatusInternalServerError
		msg := "Error al crear el To Do"
		if errors.Is(err, services.ErrConflict) {
			status = http.StatusConflict
			msg = "El To Do ya existe"
		}
		c.JSON(status, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

// GetTodoByID obtiene un To Do por ID
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	todo, err := h.todoService.GetTodoByID(id)
	if err != nil {
		status := http.StatusInternalServerError
		msg := "Error al obtener el To Do"
		if errors.Is(err, services.ErrNotFound) {
			status = http.StatusNotFound
			msg = "To Do no encontrado"
		}
		c.JSON(status, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// UpdateTodo actualiza un To Do existente
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}
	if err := h.todoService.UpdateTodo(id, updates); err != nil {
		status := http.StatusInternalServerError
		msg := "Error al actualizar el To Do"
		if errors.Is(err, services.ErrNotFound) {
			status = http.StatusNotFound
			msg = "To Do no encontrado"
		}
		c.JSON(status, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "To Do actualizado correctamente"})
}

// DeleteTodo elimina un To Do
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := h.todoService.DeleteTodo(id); err != nil {
		status := http.StatusInternalServerError
		msg := "Error al eliminar el To Do"
		if errors.Is(err, services.ErrNotFound) {
			status = http.StatusNotFound
			msg = "To Do no encontrado"
		}
		c.JSON(status, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "To Do eliminado correctamente"})
}
