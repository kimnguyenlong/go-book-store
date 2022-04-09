package graph

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContext")
	if ginContext == nil {
		err := fmt.Errorf("Could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}

type Auth struct {
	UID   string
	Email string
	Role  string
}

func GetAuthFromContext(ctx context.Context) (*Auth, error) {
	ginContext, err := GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	authorization := ginContext.GetHeader("Authorization")
	bearerTokens := strings.Split(authorization, " ")
	if len(bearerTokens) != 2 || bearerTokens[0] != "Bearer" {
		return nil, fmt.Errorf("Invalid token, please provide a Bearer token")
	}
	token, err := jwt.Parse(bearerTokens[1], func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Name {
			return nil, fmt.Errorf("Invalid signing method: %v", t.Method.Alg())
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); token.Valid && ok {
		exp := claims["exp"].(float64)
		if int64(exp) < time.Now().Unix() {
			return nil, fmt.Errorf("Token is expired")
		}
		uid := claims["uid"].(string)
		email := claims["email"].(string)
		role := claims["role"].(string)
		return &Auth{
			UID:   uid,
			Email: email,
			Role:  role,
		}, nil
	} else {
		return nil, fmt.Errorf("Invalid token")
	}
}
