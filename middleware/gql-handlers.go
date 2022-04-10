package middleware

import (
	"book-store/graph/generated"
	"book-store/graph/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GraphqlHandler(db *mongo.Database) gin.HandlerFunc {
	resolver := &resolver.Resolver{DB: db}
	schema := generated.NewExecutableSchema(generated.Config{Resolvers: resolver})
	h := handler.NewDefaultServer(schema)
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("Book Store - GraphQL", "/gql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
