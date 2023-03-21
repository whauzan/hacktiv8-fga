package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	_routes "github.com/whauzan/hacktiv8-fga/pkg/routes"
)

func main() {
	fmt.Println("Hello! Wahyu here ^-^")

	router := gin.Default()

	routesInit := _routes.HandlerList{}

	routesInit.RouteRegister(router)
}