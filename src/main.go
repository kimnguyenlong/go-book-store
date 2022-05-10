package main

import (
	"book-store/db"
	"book-store/middleware"
	"context"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	mongoClient := db.Connect(os.Getenv("MONGODB_CONNECTTION_URI"))
	defer mongoClient.Disconnect(context.Background())

	router := gin.Default()

	router.Use(middleware.GinContextToGQLContext())

	router.POST("/gql", middleware.GraphqlHandler(mongoClient.Database("book-store")))
	router.GET("/", middleware.PlaygroundHandler())

	router.Run(":9090")
}
