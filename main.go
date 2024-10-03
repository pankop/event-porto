package main

import (
	"log"
	"os"

	"github.com/pankop/event-porto/cmd"
	"github.com/pankop/event-porto/config"
	"github.com/pankop/event-porto/controller"
	"github.com/pankop/event-porto/middleware"
	"github.com/pankop/event-porto/repository"
	"github.com/pankop/event-porto/routes"
	"github.com/pankop/event-porto/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetUpDatabaseConnection()
	defer config.CloseDatabaseConnection(db)

	if len(os.Args) > 1 {
		cmd.Commands(db)
		return
	}

	var (
		jwtService service.JWTService = service.NewJWTService()

		// Implementation Dependency Injection
		// Repository
		userRepository repository.UserRepository = repository.NewUserRepository(db)

		// Service
		userService service.UserService = service.NewUserService(userRepository, jwtService)

		// Controller
		userController controller.UserController = controller.NewUserController(userService)
	)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	// routes
	routes.User(server, userController, jwtService)

	server.Static("/assets", "./assets")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	if err := server.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
