package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/whauzan/hacktiv8-fga/pkg/controllers"
)

type HandlerList struct {
	
}

func (handlerList *HandlerList) RouteRegister(g *gin.Engine) {
	api := g.Group("book/api/v1")
	api.GET("/", controllers.Hello)

	g.Run(":8080")
}