package response

import (
	"time"

	"github.com/whauzan/hacktiv8-fga/business/books"
)

type BookResponse struct {
	Message   string    `json:"message"`
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Desc      string    `json:"desc"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

func FromDomainHello(domain books.Hello) MessageResponse {
	return MessageResponse{
		Message: domain.Message,
	}
}

func FromDomainInsert(domain books.Domain) BookResponse {
	return BookResponse{
		Message:   "Insert Book Success",
		ID:        domain.ID,
		Title:     domain.Title,
		Author:    domain.Author,
		Desc:      domain.Desc,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainUpdate(domain books.Domain) BookResponse {
	return BookResponse{
		Message:   "Update Book Success",
		ID:        domain.ID,
		Title:     domain.Title,
		Author:    domain.Author,
		Desc:      domain.Desc,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainDelete(domain books.Domain) MessageResponse {
	return MessageResponse{
		Message: "Delete Book Success",
	}
}
