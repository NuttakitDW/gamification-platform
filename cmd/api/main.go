package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"

	_ "github.com/NuttakitDW/gamification-platform/cmd/api/docs"
	"github.com/NuttakitDW/gamification-platform/internal/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gamification Platform API
// @version 1.0
// @description This is a api server for gamification platform.
// @BasePath /api/v1
func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	initDatabase()

	r := gin.New()

	r.GET("/v1/api/helloworld/:name", handlers.GetStringByInt)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8080")
}

func initDatabase() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to the database!")


}