package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"template-go/controllers"
	"template-go/databases"
	"template-go/helpers"
	"template-go/middlewares"
	"template-go/repositories"
	"template-go/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := databases.DatabaseInit()

	// Initialize User components
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(db, userRepository)
	userController := controllers.NewUserController(userService)

	// Initialize Auth components
	authService := services.NewAuthService(db, userRepository)
	authController := controllers.NewAuthController(authService)

	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// Protected user routes
	router.Route("/api/users", func(r chi.Router) {
		r.Use(middlewares.ProtectedRoute)
		r.Get("/", userController.FindAll)                      // Get all users
		r.Post("/", userController.Create)                      // Create user
		r.Get("/{identifier}", userController.FindByIdentifier) // Get user by identifier
		r.Delete("/{userId}", userController.Delete)            // Delete user by userId
		r.Put("/{userId}", userController.Update)               // Update user by userId
	})

	// Authentication routes (public)
	router.Post("/api/auth/register", authController.Register)
	router.Post("/api/auth/login", authController.Login)

	// Start the server
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler: router,
	}

	log.Println("Listening on port", server.Addr)

	err = server.ListenAndServe()
	helpers.PanicOnError(err)
}
