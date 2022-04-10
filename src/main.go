package main

import (
	"book-store/db"
	"book-store/middleware"
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("MONGODB_CONNECTTION_URI") == "" || os.Getenv("JWT_LIFE_TIME") == "" || os.Getenv("JWT_SECRET") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalf("Error when loading .env file: %v", err.Error())
		}
	}

	mongoClient := db.Connect(os.Getenv("MONGODB_CONNECTTION_URI"))
	defer mongoClient.Disconnect(context.Background())

	router := gin.Default()

	router.Use(middleware.GinContextToGQLContext())

	router.POST("/gql", middleware.GraphqlHandler(mongoClient.Database("book-store")))
	router.GET("/", middleware.PlaygroundHandler())

	router.Run(":8080")
}
