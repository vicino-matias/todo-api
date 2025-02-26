package routes

import (
	"log"
	"net/http"
	"todo-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

// healthCheckHandler responde con un estado OK
func healthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// loggingMiddleware registra cada solicitud HTTP
func loggingMiddleware(c *gin.Context) {
	log.Printf("Solicitud recibida: %s %s", c.Request.Method, c.Request.URL.Path)
	c.Next()
}

// RegisterRoutes configura las rutas de la API
func RegisterRoutes(r *gin.Engine, todoHandler *handlers.TodoHandler) {
	// Ruta para verificar el estado del servicio
	r.GET("/health", healthCheckHandler)

	// Grupo de rutas de la API v1
	v1 := r.Group("/api/v1", loggingMiddleware)

	// Rutas CRUD para To Do's
	v1.GET("/todos", todoHandler.GetTodos)          // Obtener todos los To Do's
	v1.POST("/todos", todoHandler.CreateTodo)       // Crear un nuevo To Do
	v1.GET("/todos/:id", todoHandler.GetTodoByID)   // Obtener un To Do por ID
	v1.PUT("/todos/:id", todoHandler.UpdateTodo)    // Actualizar un To Do
	v1.DELETE("/todos/:id", todoHandler.DeleteTodo) // Eliminar un To Do por ID
}
