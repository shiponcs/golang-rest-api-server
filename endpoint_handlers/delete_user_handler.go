package endpoint_handlers

import (
	"fmt"
	"github.com/bookstore-rest-api-server/db"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var id int
	id, err := strconv.Atoi(chi.URLParam(request, "id"))
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if _, ok := db.Books[id]; !ok {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("not found"))
		return
	}
	delete(db.Books, id)
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Deleted"))
}
