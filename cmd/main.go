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
		log.Fatalf("❌Error al cargar las configuraciones: %v", err)
	}

	// Inicializar la base de datos
	db := initDatabase(cfg)

	todoRepo := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(*todoService)

	// Inicializar servidor HTTP
	r := gin.Default()

	// Middleware para manejo de errores
	r.Use(errorHandlerMiddleware)

	// Registrar rutas de la API

	routes.RegisterRoutes(r, todoHandler)

	// Iniciar servidor
	log.Printf("🚀 Servidor corriendo en http://localhost:%s", cfg.ServerPort)
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("❌ Error al iniciar el servidor: %v", err)
	}
}

// Funcion para inicializar la base de datos
func initDatabase(cfg *configs.Config) *gorm.DB {
	dsn := buildDSN(cfg)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌Error al conectar la base de datos: %v", err)
	}

	//Migrar el modelos
	if err := db.AutoMigrate(&models.Todo{}); err != nil {
		log.Fatalf("❌Error al migrar la base de datos: %v", err)
	}

	log.Println("✅ Base de datos conectada y migrada correctamente")
	return db
}

// Funcion middleware centralizado para manejo de errores
func errorHandlerMiddleware(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": c.Errors.JSON()})
	}
}

// Construye la cadena de conexiones (DSN) para PostgreSQL
func buildDSN(cfg *configs.Config) string {
	return "host=" + cfg.DBHost +
		" user=" + cfg.DBUser +
		" password=" + cfg.DBPassword +
		" dbname=" + cfg.DBName +
		" port=" + cfg.DBPort +
		" sslmode=disable TimeZone=UTC"
}
