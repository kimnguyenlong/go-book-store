package main

import (
	"book-store/db"
	"book-store/handler"
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error when loading .env file: %v", err.Error())
	}

	mongoClient := db.Connect(os.Getenv("MONGODB_CONNECTTION_URI"))
	defer mongoClient.Disconnect(context.Background())

	router := gin.Default()

	router.POST("/gql", handler.GraphqlHandler(mongoClient.Database("book-store")))
	router.GET("/", handler.PlaygroundHandler())

	router.Run(":8080")
}
