package books

import (
	"time"

	"github.com/whauzan/hacktiv8-fga/business"
)

type BookService struct {
	repository     Repository
	contextTimeout time.Duration
}

func NewBookService(repo Repository, timeout time.Duration) Service {
	return &BookService{
		repository:     repo,
		contextTimeout: timeout,
	}
}

func (service *BookService) Hello() (Hello, error) {
	books := Hello{
		Message: "Welcome To CRUDBook",
	}

	return books, nil
}

func (service *BookService) GetAllBook() ([]Domain, error) {
	books, _ := service.repository.GetAllBook()

	// if len(books) == 0 {
	// 	return nil, business.ErrNotFound
	// }

	return books, nil
}

func (service *BookService) GetBookByID(id int) (Domain, error) {
	book, err := service.repository.GetBookByID(id)

	if err != nil {
		return Domain{}, business.ErrNotFound
	}

	return book, nil
}

func (service *BookService) AddBook(domain *Domain) (Domain, error) {
	if domain.Title == "" {
		return Domain{}, business.ErrEmptyForm
	}
	if domain.Author == "" {
		return Domain{}, business.ErrEmptyForm
	}
	if domain.Desc == "" {
		return Domain{}, business.ErrEmptyForm
	}

	book, err := service.repository.Insert(domain)

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	
	return book, nil
}

func (service *BookService) UpdateBook(id int, domain *Domain) (Domain, error) {
	book, err := service.repository.Update(id, domain)

	if err != nil {
		return Domain{}, business.ErrInternalServer
	}
	return book, nil
}

func (service *BookService) DeleteBook(id int) (Domain, error) {
	book, err := service.repository.Delete(id)
	
	if err != nil {
		return Domain{}, business.ErrInternalServer
	}

	return book, nil
}