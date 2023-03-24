package books

import (
	"github.com/whauzan/hacktiv8-fga/business/books"
	"gorm.io/gorm"
)

type PostgreSQLBookRepository struct {
	Conn *gorm.DB
}

func NewPostgreSQLBookRepository(conn *gorm.DB) books.Repository {
	return &PostgreSQLBookRepository{
		Conn: conn,
	}
}

func (rep *PostgreSQLBookRepository) GetAllBook() ([]books.Domain, error) {
	var book []Books
	result := rep.Conn.Find(&book)

	if result.Error != nil {
		return nil, result.Error
	}

	return toListDomain(book), nil
}

func (rep *PostgreSQLBookRepository) GetBookByID(id int) (books.Domain, error) {
	var book Books
	result := rep.Conn.First(&book, "ID = ?", id)

	if result.Error != nil {
		return books.Domain{}, result.Error
	}

	return toDomain(book), nil
}

func (rep *PostgreSQLBookRepository) Insert(domain *books.Domain) (books.Domain, error) {
	book := fromDomain(*domain)

	result := rep.Conn.Create(&book)

	if result.Error != nil {
		return books.Domain{}, result.Error
	}

	return toDomain(book), nil
}

func (rep *PostgreSQLBookRepository) Update(id int, domain *books.Domain) (books.Domain, error) {
	var book Books
	bookData := fromDomain(*domain)

	result := rep.Conn.First(&book, "ID = ?", id)

	if result.Error != nil {
		return books.Domain{}, result.Error
	}

	update := rep.Conn.Where("id", id).Updates(&bookData)
	if update.Error != nil {
		return books.Domain{}, update.Error
	}

	return toDomain(bookData), nil
}

func (rep *PostgreSQLBookRepository) Delete(id int) (books.Domain, error) {
	var book Books

	result := rep.Conn.Delete(&book, id)

	if result.Error != nil {
		return books.Domain{}, result.Error
	}

	return toDomain(book), nil
}