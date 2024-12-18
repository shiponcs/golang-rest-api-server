package endpoint_handlers

import (
	"encoding/json"
	"github.com/bookstore-rest-api-server/db"
	"net/http"
)

func GetUsers(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	booksSlice := make([]map[string]interface{}, 0, len(db.Books))
	for id, book := range db.Books {
		booksSlice = append(booksSlice, map[string]interface{}{
			"id":       id,
			"author":   book.Auth,
			"title":    book.Title,
			"category": book.Category,
			"isbn":     book.ISBN,
		})
	}
	if err := json.NewEncoder(writer).Encode(booksSlice); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.WriteHeader(http.StatusOK)
}
