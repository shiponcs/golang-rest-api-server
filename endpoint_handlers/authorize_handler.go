package endpoint_handlers

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		fmt.Println("Error: SECRET_KEY environment variable is not set")
		return
	}

	claims := jwt.MapClaims{
		"name": os.Getenv("USER_NAME"),
		"exp":  time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Printf("Error creating token: %v\n", err)
		return
	}
	_, err = w.Write([]byte("token " + tokenString))
	if err != nil {
		return
	}
	fmt.Println("Generated JWT Token:", tokenString)
}
