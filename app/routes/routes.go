package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/whauzan/hacktiv8-fga/app/handler/book"
)

type HandlerList struct {
	BookHandler book.BookHandler
}

func (handlerList *HandlerList) RouteRegister(g *gin.Engine) {
	api := g.Group("book/api/v1")
	api.GET("/", handlerList.BookHandler.Hello)

	books := api.Group("/books")
	books.GET("/", handlerList.BookHandler.GetAllBook)
	books.GET("/:id", handlerList.BookHandler.GetBookByID)
	books.POST("/", handlerList.BookHandler.AddBook)
	books.PUT("/:id", handlerList.BookHandler.UpdateBook)
	books.DELETE("/:id", handlerList.BookHandler.DeleteBook)

	g.Run(viper.GetString("server.address"))
}