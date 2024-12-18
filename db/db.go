package db

import "github.com/bookstore-rest-api-server/model"

var Books map[int]model.Book

func Init() {
	Books = make(map[int]model.Book)
	Books = map[int]model.Book{
		0: {
			Auth:     "The MANIAC",
			Title:    "Benjamín Labatut",
			Category: "Novel",
			ISBN:     "9780593654491",
		},
		1: {
			Auth:     "George Orwell",
			Title:    "1984",
			Category: "Dystopian",
			ISBN:     "9780451524935",
		},
		2: {
			Auth:     "J.K. Rowling",
			Title:    "Harry Potter and the Sorcerer's Stone",
			Category: "Fantasy",
			ISBN:     "9780439708180",
		},
		3: {
			Auth:     "Yuval Noah Harari",
			Title:    "Sapiens: A Brief History of Humankind",
			Category: "Non-Fiction",
			ISBN:     "9780062316097",
		},
		4: {
			Auth:     "Michelle Obama",
			Title:    "Becoming",
			Category: "Biography",
			ISBN:     "9781524763138",
		},
	}

}
