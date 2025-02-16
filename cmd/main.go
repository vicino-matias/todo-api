package main

import (
	"log"
	"net/http"
	"todo-api/configs"
	"todo-api/internal/handlers"
	"todo-api/internal/models"
	"todo-api/internal/repositories"
	"todo-api/internal/routes"
	"todo-api/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Cargamos configuraciones
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Error al cargar las configuraciones: %v", err)
	}
	// Inicializar la base de datos
	dsn := buildDSN(cfg)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error al conectar la base de datos: %v", err)
	}

	//Migrar el modelo de To Do (crear la tabla si no existe)
	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatalf("Error al migrar la base de datos: %v", err)
	}

	// Inicializar el repositorio
	todoRepo := repositories.NewTodoRepository(db)

	// Inicializar servicio
	todoService := services.NewTodoService(*todoRepo)

	// Inicializar los handlers
	todoHandler := handlers.NewTodoHandler(*todoService)

	//Inicializar servidor
	r := gin.Default()

	// Middleware para manejo de errores centralizados
	r.Use(func(c *gin.Context) {
		c.Next()

		// Verificar si hubo algun error durante el proceso
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": c.Errors.JSON(),
			})
		}
	})

	// Registra rutas
	routes.RegisterRoutes(r, todoHandler)

	//Correr servidor
	log.Printf("Servidor corriendo en http://localhost:%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

// Construye la cadena de conexion (DSN) para PostgreSQL
func buildDSN(cfg *configs.Config) string {
	return "host=" + cfg.DBHost +
		" user=" + cfg.DBUser +
		" password=" + cfg.DBPassword +
		" dbname=" + cfg.DBName +
		" port=" + cfg.DBPort +
		" sslmode=disable TimeZone=UTC"
}
