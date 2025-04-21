package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nedokyrill/auth-service/internal/handlers/authHandler"
	"github.com/nedokyrill/auth-service/internal/repository/authRepository"
	"github.com/nedokyrill/auth-service/internal/repository/userRepository"
	"github.com/nedokyrill/auth-service/internal/service/authService"
	"github.com/nedokyrill/auth-service/pkg/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func Run() {
	if err := godotenv.Load(); err != nil { //в будущем можно вытаскивать из докера
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	log.Println("Successfully connected to the database")

	//repo
	authRepo := authRepository.NewAuthRepository(db)
	userRepo := userRepository.NewUserRepository(db)

	//service
	authServ := authService.NewAuthService(authRepo, userRepo)

	//handler
	authHand := authHandler.NewAuthHandler(authServ)

	//default route
	router := gin.Default()
	api := router.Group("/api/v1/")

	//REGISTER_ROUTES
	authHand.RegisterRoutes(api)

	server := utils.NewAPIServer(os.Getenv("ADDR"), router)
	log.Println("Server success running on port: ", os.Getenv("ADDR"))
	utils.Start(server)
}
