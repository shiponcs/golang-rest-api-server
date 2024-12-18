package endpoint_handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bookstore-rest-api-server/db"
	"github.com/bookstore-rest-api-server/model"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func UpdateUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var id int
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		fmt.Println(err)
	}
	book, ok := db.Books[id]
	if !ok {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("not found"))
		return
	}

	var sentBook model.Book
	err = json.NewDecoder(request.Body).Decode(&sentBook)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		writer.Write([]byte(err.Error()))
		return
	}
	db.Books[id] = sentBook
	book = db.Books[id]
	bookMap := map[string]interface{}{
		"id":       id,
		"author":   book.Auth,
		"title":    book.Title,
		"category": book.Category,
		"isbn":     book.ISBN,
	}
	if err = json.NewEncoder(writer).Encode(bookMap); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	writer.WriteHeader(http.StatusOK)
}
