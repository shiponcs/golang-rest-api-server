# A REST API Server in Golang

## Dependecies
- [go](https://github.com/golang)
- [spf13/cobra](https://github.com/spf13/cobra),
- [go-chi/chi](https://github.com/go-chi/chi), 
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt)

## Build and Run the project
```bash
  > make all
  mkdir -p bin
  go build -o bin/book-store-api-server
```

```bash
  ❯ bin/book-store-api-server serve --port 8080                                                                                                                                         main ✱ ◼
  serving at port 8080
```


## API Endpoints
- POST /authorize
    - Implements Basic Auth
  ```bash
    ❯ curl -X POST --user admin:admin  localhost:8080/authorize
    token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzQ1NTA4ODQsIm5hbWUiOiJhZG1pbiJ9.7eyf8XmnZHgG0CNUQXJqjy4m8cvGITCkdUg2eFJ1CLo
    ```  
- GET /books
    ```bash
      ❯ curl -X GET -H "Authorization: Bearer <token>" localhost:8080/books
      [{"author":"The MANIAC","category":"Novel","id":0,"isbn":"9780593654491","title":"Benjamín Labatut"},{"author":"George Orwell","category":"Dystopian","id":1,"isbn":"9780451524935","title":"1984"},{"author":"J.K. Rowling","category":"Fantasy","id":2,"isbn":"9780439708180","title":"Harry Potter and the Sorcerer's Stone"},{"author":"Yuval Noah Harari","category":"Non-Fiction","id":3,"isbn":"9780062316097","title":"Sapiens: A Brief History of Humankind"},{"author":"Michelle Obama","category":"Biography","id":4,"isbn":"9781524763138","title":"Becoming"}]

   ```
- GET /book/{id}
    ```bash
  ❯ curl -X GET -H "Authorization: Bearer <token> localhost:8080/book/2
    {"author":"J.K. Rowling","category":"Fantasy","id":2,"isbn":"9780439708180","title":"Harry Potter and the Sorcerer's Stone"}
   ```
- DELETE /book/{id}
    ```bash
  ❯ curl -X DELETE -H "Authorization: Bearer <token> localhost:8080/book/1
    Deleted
   ```
- PUT /book/{id}
    ```bash
    ❯ curl -X PUT -H "Content-Type: application/json" \
      -H "Authorization: Bearer <token>" \
      -d '{
      "author": "J.K. Rowling",
      "title": "Harry Potter and the Chamber of Secrets",
     "category": "Fantasy",
     "isbn": "9780439064873"
      }'\
      http://localhost:8080/book/3
    { "author":"","category":"Fantasy","id":3,"isbn":"9780439064873","title":"Harry Potter and the Chamber of Secrets"}
   ```
## Clean
```bash
  > make clean
  rm -rf bin
```

## Docker 
### Build Docker Image
```bash
> make docker-build
```
### Run Docker Image
#### Run with default options
```bash
> docker run -p 8080:8080 book-store-api-server
serving at port 8080
```
Default port is 8080 and other default values can be found at .env file.

#### Pass Command Line Arg values
##### Modify the listening port
```bash
> docker run -p 8081:8081 book-store-api-server ./book-store-api-server serve --port 8081
serving at port 8081

``` 
#### Modify environment values
##### Modify the password
```bash
❯ docker run -p 8080:8080 --env PASSWORD=admin book-store-api-server
serving at port 8080
```
##### Override the .env file
###### Mount the new .env file
```bash
❯ docker run -p 8080:8080 -v $(pwd)/.env:/app/.env book-store-api-server
serving at port 8080
```

