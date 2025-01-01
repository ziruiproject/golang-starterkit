package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"technical-test-go/controllers"
	"technical-test-go/helpers"
	"technical-test-go/middlewares"
	"technical-test-go/repositories"
	"technical-test-go/services"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname),
	)
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}

	// Initialize User components
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(db, userRepository)
	userController := controllers.NewUserController(userService)

	// Initialize Auth components
	authService := services.NewAuthService(db, userRepository)
	authController := controllers.NewAuthController(authService)

	// Initialize Product components
	productRepository := repositories.NewProductRepository() // Assuming the repository constructor takes *sqlx.DB
	productService := services.NewProductService(db, productRepository, userRepository)
	productController := controllers.NewProductController(productService)

	// Initialize Product components
	bankRepository := repositories.NewBankRepository() // Assuming the repository constructor takes *sqlx.DB
	bankService := services.NewBankService(db, bankRepository, userRepository)
	bankController := controllers.NewBankController(bankService)

	orderRepository := repositories.NewOrderRepository()

	// Initialize Product components
	cartRepository := repositories.NewCartRepository() // Assuming the repository constructor takes *sqlx.DB
	cartService := services.NewCartService(db, cartRepository, orderRepository, productService, userService, bankService)
	cartController := controllers.NewCartController(cartService)

	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)    // Log HTTP requests
	router.Use(middleware.Recoverer) // Recover from panics
	// User routes with ProtectedRoute middleware
	router.Route("/api/users", func(r chi.Router) {
		r.Use(middlewares.ProtectedRoute)  // Apply JWTMiddleware to protect all user routes
		r.Get("/", userController.FindAll) // Get all users
		r.Post("/", userController.Create)
		r.Get("/{identifier}", userController.FindByIdentifier)     // Get user by identifier
		r.Delete("/{userId}", userController.Delete)                // Delete user by userId
		r.Put("/{userId}", userController.Update)                   // Update user by userId
		r.Get("/{userId}/products", productController.FindByUserId) // Find products by userId
	})

	// Authentication routes (public)
	router.Post("/api/auth/register", authController.Register)
	router.Post("/api/auth/login", authController.Login)

	// Public product routes
	router.Get("/api/products/{productId}", productController.FindById) // Find product by productId
	router.Get("/api/search", productController.FindBySearch)           // Search route using query parameter

	// Protected product routes (wrapped with middleware)
	router.Route("/api/products", func(r chi.Router) {
		r.Use(middlewares.ProtectedRoute)                  // Apply JWTMiddleware to protect product routes
		r.Get("/", productController.FindAll)              // Get all products
		r.Post("/", productController.Create)              // Create a product
		r.Put("/{productId}", productController.Update)    // Update a product
		r.Delete("/{productId}", productController.Delete) // Delete a product
	})

	// Protected cart routes
	router.Route("/api/carts", func(r chi.Router) {
		r.Use(middlewares.ProtectedRoute)                          // Apply JWTMiddleware to protect cart routes
		r.Get("/", cartController.FindAll)                         // Get all cart items
		r.Post("/", cartController.Create)                         // Add item to cart
		r.Put("/{cartId}", cartController.Update)                  // Update cart item by cartId
		r.Delete("/{cartId}", cartController.Delete)               // Delete cart item by cartId
		r.Get("/user/{userId}", cartController.FindByUserId)       // Get cart items by userId
		r.Post("/user/{userId}/checkout", cartController.Checkout) // Get cart items by userId
		r.Get("/{cartId}", cartController.FindById)                // Get cart item by cartId
	})

	// Protected bank routes
	router.Route("/api/banks", func(r chi.Router) {
		r.Use(middlewares.ProtectedRoute)               // Apply JWTMiddleware to protect bank routes
		r.Post("/", bankController.CreateAccount)       // Create a new bank account
		r.Get("/", bankController.GetAllAccounts)       // Get all bank accounts
		r.Get("/{id}", bankController.GetAccountById)   // Get bank account by ID
		r.Put("/{id}", bankController.UpdateAccount)    // Update bank account by ID
		r.Delete("/{id}", bankController.DeleteAccount) // Delete bank account by ID
		r.Post("/transfer", bankController.Transfer)    // Transfer between accounts
	})

	// Start the server
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Listening on port", server.Addr)

	err = server.ListenAndServe()
	helpers.PanicOnError(err)
}
