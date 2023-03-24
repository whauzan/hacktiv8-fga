package repository

import (
	"gorm.io/gorm"
	
	bookDomain "github.com/whauzan/hacktiv8-fga/business/books"
	bookDB "github.com/whauzan/hacktiv8-fga/repository/databases/books"
)

func NewBookRepository(conn *gorm.DB) bookDomain.Repository {
	return bookDB.NewPostgreSQLBookRepository(conn)
}