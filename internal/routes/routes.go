package routes

import (
	"log"
	"net/http"
	"todo-api/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, todoHandler *handlers.TodoHandler) {

	// Ruta para health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Grupo de rutas para la API
	v1 := r.Group("/api/v1")
	{
		// Middleware de logging
		v1.Use(func(c *gin.Context) {
			log.Printf("Solicitud recibida: %s %s", c.Request.Method, c.Request.URL.Path)
			c.Next()
		})

		// Rutas de la API

		//Obtener todos los To Do's
		v1.GET("/todos", todoHandler.GetTodos)

		// Crear un nuevo To Do
		v1.POST("/todos", todoHandler.CreateTodo)

		// Obtener un To Do por su ID
		v1.GET("/todos/:id", todoHandler.GetTodoByID)

		// Actualizar un To Do existente
		v1.PUT("/todos/:id", todoHandler.UpdateTodo)

		// Eliminar un To Do por su ID
		v1.DELETE("/todos/:id", todoHandler.DeleteTodo)
	}
}
