package models

type Book struct {
	Id     int    `json:"id" gorm:"primary key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
	Pages  int    `json:"pages"`
}
