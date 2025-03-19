package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/ggicci/httpin"
	httpinintegration "github.com/ggicci/httpin/integration"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"github.com/vektah/gqlparser/v2/ast"
	"net/http"
	"os"
	"template-go/commons/config"
	"template-go/controllers"
	"template-go/databases"
	"template-go/graph"
	"template-go/helpers"
	"template-go/middlewares"
	"template-go/models/web"
	"template-go/repositories"
	"template-go/services"
	"time"
)

func init() {
	// Register a directive named "path" to retrieve values from `chi.URLParam`,
	// i.e. decode path variables.
	httpinintegration.UseGochiURLParam("path", chi.URLParam)
}

func main() {
	// Initialize Logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	err := godotenv.Load()
	if err != nil {
		log.Fatal().Msg("Error loading .env file")
	}

	// Initialize Config
	config := config.NewConfig()

	// Initialize Database
	db := databases.DatabaseInit()

	// Initialize User components
	userRepository := repositories.NewUserRepository()
	userService := services.NewUserService(db, userRepository)
	userController := controllers.NewUserController(userService, config.Sanitation)

	// Initialize Auth components
	authService := services.NewAuthService(db, userRepository)
	authController := controllers.NewAuthController(authService)

	// Router
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Recoverer)
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(hlog.NewHandler(log.Logger))
	router.Use(hlog.AccessHandler(func(r *http.Request, status, size int, duration time.Duration) {
		log.Info().
			Str("method", r.Method).
			Str("path", r.URL.Path).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("HTTP request")
	}))

	// Set Timeout
	router.Use(middleware.Timeout(60 * time.Second))

	// Initialize Default Params
	router.With(httpin.NewInput(web.DefaultParams{}))

	// GraphQL
	graphql := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	graphql.AddTransport(transport.Options{})
	graphql.AddTransport(transport.GET{})
	graphql.AddTransport(transport.POST{})

	graphql.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	graphql.Use(extension.Introspection{})
	graphql.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	router.Route("/api/v1", func(r chi.Router) {
		// Protected user routes
		r.Route("/users", func(r chi.Router) {
			r.Use(middlewares.ProtectedRoute)
			r.Get("/", userController.FindAll)                      // Get all users
			r.Post("/", userController.Create)                      // Create user
			r.Get("/{identifier}", userController.FindByIdentifier) // Get user by identifier
			r.Delete("/{userId}", userController.Delete)            // Delete user by userId
			r.Put("/{userId}", userController.Update)               // Update user by userId
		})

		// Authentication routes (public)
		r.Post("/auth/register", authController.Register)
		r.Post("/auth/login", authController.Login)

		r.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
		r.Handle("/graphql", graphql)
	})

	// Start the server
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler: router,
	}

	chi.Walk(router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Info().
			Str("route", route).
			Str("method", method).
			Int("middlewares", len(middlewares)).
			Msg("Registered Route")
		return nil
	})

	log.Info().Msgf("Listening on port %s", server.Addr)

	err = server.ListenAndServe()
	helpers.PanicOnError(err)
}
