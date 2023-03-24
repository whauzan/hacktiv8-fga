package request

import "github.com/whauzan/hacktiv8-fga/business/books"

type CreateBook struct {
}

type Books struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

func (req *Books) ToDomain() *books.Domain {
	return &books.Domain{
		Title:  req.Title,
		Author: req.Author,
		Desc:   req.Desc,
	}
}
