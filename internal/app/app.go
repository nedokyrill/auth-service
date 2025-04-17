package app

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
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

	_, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error with connection to database: ", err)
	}
	log.Println("Connected to database successfully", os.Getenv("DATABASE_URL"))

	//repo

	//service

	//handler
	//
	////default route
	//router := gin.Default()
	//api := router.Group("/api/v1/")
	//
	////REGISTER_ROUTES
	//
	//server := Utils.NewServer(os.Getenv("ADDR"), router)
	//log.Println("Server success running on port: ", os.Getenv("ADDR"))
	//Utils.Start(server)
}
