package model

import (
	"book-store/utils"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func (user User) CreateJWT() (string, error) {
	jwtLifeTime, err := strconv.Atoi(os.Getenv("JWT_LIFE_TIME"))
	if err != nil {
		jwtLifeTime = utils.DEFAULT_JWT_LIFE_TIME
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid":   user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24 * time.Duration(jwtLifeTime)).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (user User) CheckPassword(candidatePassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(candidatePassword)) == nil
}
