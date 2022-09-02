package mocks

import "helo/models"

//import "github.com/tutorials/go/crud/pkg/models"
//import "github.com/ThotPrime/Project/tree/main/Project/pkg/models"

var Books = []models.Book{
	{
		Id:     1,
		Title:  "Golang",
		Author: "Gopher",
		Desc:   "A book for Go",
		Pages:  27,
	},
}
