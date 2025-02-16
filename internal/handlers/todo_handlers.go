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

// GetTodos obtiene todos los To Do's existentes.
// @Summary Obtener todos los To Do's
// @Description Obtiene una lista de todos los To Do's existentes.
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} models.Todo
// @Failure 500 {object} map[string]string
// @Router /todos [get]

func (h *TodoHandler) GetTodos(c *gin.Context) {
	todos, err := h.todoService.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// CreateTodo crea un nuevo To Do.
// @Summary Crear un nuevo To Do
// @Description Crea un nuevo To Do con los datos proporcionados.
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body models.Todo true "Datos del To Do"
// @Success 201 {object} models.Todo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos [post]

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validacion manual
	if todo.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El campo 'title' es requerido"})
		return
	}

	if err := h.todoService.CreateTodo(&todo); err != nil {
		if errors.Is(err, services.ErrConflict) {
			c.JSON(http.StatusConflict, gin.H{"error": "El To Do ya existe"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// GetTodoByID obtiene un To Do por su ID.
// @Summary Obtener un To Do por ID
// @Description Obtiene un To Do específico por su ID.
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "ID del To Do"
// @Success 200 {object} models.Todo
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/{id} [get]

func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	id := c.Param("id")
	todo, err := h.todoService.GetTodoByID(id)
	if err != nil {
		if errors.Is(err, services.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "To Do no encontrado"})
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// UpdateTodo actualiza un To Do existente.
// @Summary Actualizar un To Do
// @Description Actualiza un To Do existente con los datos proporcionados.
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "ID del To Do"
// @Param updates body map[string]interface{} true "Campos a actualizar"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/{id} [put]

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.todoService.UpdateTodo(id, updates); err != nil {
		if errors.Is(err, services.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo no encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo actualizado correctamente"})
}

// DeleteTodo elimina un To Do por su ID.
// @Summary Eliminar un To Do
// @Description Elimina un To Do existente por su ID.
// @Tags todos
// @Accept json
// @Produce json
// @Param id path string true "ID del To Do"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/{id} [delete]

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := h.todoService.DeleteTodo(id); err != nil {
		if errors.Is(err, services.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "To Do no encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "To do eliminado correctamente"})
}
