package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println(authHeader)
		authTypeValue := strings.Split(authHeader, " ")
		if len(authTypeValue) == 2 && authTypeValue[0] == "Basic" {
			username, password, ok := r.BasicAuth()
			fmt.Println(username, password, ok)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("Invalid username or password"))
				if err != nil {
					return
				}
				return
			}
			if os.Getenv("USERNAME") != username || os.Getenv("PASSWORD") != password {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write([]byte("Invalid username or password"))
				if err != nil {
					return
				}
				return
			}
		} else if authTypeValue[0] == "Bearer" && len(authTypeValue) == 2 {
			fmt.Println("Bearer Token", authTypeValue[1])
		}
		next.ServeHTTP(w, r)
	})
}
