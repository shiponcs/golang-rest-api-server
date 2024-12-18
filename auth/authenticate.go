package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		authTypeValue := strings.Split(authHeader, " ")
		if len(authTypeValue) == 2 && authTypeValue[0] == "Basic" {
			username, password, ok := r.BasicAuth()
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("Invalid username or password"))
				if err != nil {
					return
				}
				return
			}
			if os.Getenv("USER_NAME") != username || os.Getenv("PASSWORD") != password {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("Invalid username or password"))
				if err != nil {
					return
				}
				return
			}
		} else if authTypeValue[0] == "Bearer" && len(authTypeValue) == 2 {
			if err := ParseJWT(authTypeValue[1]); err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Invalid token"))
				return
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("Invalid authorization header"))
			if err != nil {
				return
			}
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ParseJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println("user ", claims["name"], " authorized")
	} else {
		return err
	}
	return nil
}
