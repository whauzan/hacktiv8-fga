package books

import "time"

type Hello struct {
	Message string
}

type Domain struct {
	ID        int
	Title     string
	Author    string
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Hello() (Hello, error)
	GetAllBook() ([]Domain, error)
	GetBookByID(id int) (Domain, error)
	AddBook(book *Domain) (Domain, error)
	UpdateBook(id int, book *Domain) (Domain, error)
	DeleteBook(id int) (Domain, error)
}

type Repository interface {
	GetAllBook() ([]Domain, error)
	GetBookByID(id int) (Domain, error)
	Insert(book *Domain) (Domain, error)
	Update(id int, book *Domain) (Domain, error)
	Delete(id int) (Domain, error)
}
