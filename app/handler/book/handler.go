package book

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_handler "github.com/whauzan/hacktiv8-fga/app/handler"
	"github.com/whauzan/hacktiv8-fga/app/handler/book/request"
	"github.com/whauzan/hacktiv8-fga/app/handler/book/response"
	"github.com/whauzan/hacktiv8-fga/business/books"
)

type BookHandler struct {
	BookService books.Service
}

func NewBookHandler(service books.Service) *BookHandler {
	return &BookHandler{
		BookService: service,
	}
}

func (handler *BookHandler) Hello(ctx *gin.Context) {
	data, _ := handler.BookService.Hello()

	_handler.NewSuccessResponse(ctx, response.FromDomainHello(data))
}

func (handler *BookHandler) GetAllBook(ctx *gin.Context) {
	data, err := handler.BookService.GetAllBook()

	if err != nil {
		_handler.NewErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	_handler.NewSuccessResponse(ctx, data)
}

func (handler *BookHandler) GetBookByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := handler.BookService.GetBookByID(id)

	if err != nil {
		_handler.NewErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	_handler.NewSuccessResponse(ctx, data)
}

func (handler *BookHandler) AddBook(ctx *gin.Context) {
	req := request.Books{}

	if err := ctx.Bind(&req); err != nil {
		_handler.NewErrorResponse(ctx, http.StatusBadRequest, err)
	}

	data, err := handler.BookService.AddBook(req.ToDomain())

	if err != nil {
		_handler.NewErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	_handler.NewSuccessResponse(ctx, response.FromDomainInsert(data))
}

func (handler *BookHandler) UpdateBook(ctx *gin.Context) {
	req := request.Books{}

	if err := ctx.Bind(&req); err != nil {
		_handler.NewErrorResponse(ctx, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := handler.BookService.UpdateBook(id, req.ToDomain())

	if err != nil {
		_handler.NewErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	_handler.NewSuccessResponse(ctx, response.FromDomainUpdate(data))
}

func (handler *BookHandler) DeleteBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	data, err := handler.BookService.DeleteBook(id)

	if err != nil {
		_handler.NewErrorResponse(ctx, http.StatusInternalServerError, err)
		return
	}

	_handler.NewSuccessResponse(ctx, response.FromDomainDelete(data))
}