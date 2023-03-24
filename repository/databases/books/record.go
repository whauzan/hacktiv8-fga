package books

import (
	"time"

	"github.com/whauzan/hacktiv8-fga/business/books"
)

type Books struct {
	ID        int    `gorm:"primaryKey"`
	Title     string `gorm:"not null;unique;type:varchar(191)"`
	Author    string `gorm:"not null;type:varchar(191)"`
	Desc      string `gorm:"not null;type:varchar(191)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(book Books) books.Domain {
	return books.Domain{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		Desc:      book.Desc,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
}

func toListDomain(domain []Books) (result []books.Domain) {
	result = []books.Domain{}
	for _, domain := range domain {
		result = append(result, toDomain(domain))
	}

	return result
}

func fromDomain(domain books.Domain) Books {
	return Books{
		ID:        domain.ID,
		Title:     domain.Title,
		Author:    domain.Author,
		Desc:      domain.Desc,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
