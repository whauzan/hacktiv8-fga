package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	
	_routes "github.com/whauzan/hacktiv8-fga/app/routes"
	_dbDriver "github.com/whauzan/hacktiv8-fga/repository/postgreSQL"

	_bookHandler "github.com/whauzan/hacktiv8-fga/app/handler/book"
	_bookService "github.com/whauzan/hacktiv8-fga/business/books"
	_bookRepo "github.com/whauzan/hacktiv8-fga/repository/databases/books"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_bookRepo.Books{},
	)
}

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	fmt.Println("Hello! Wahyu here ^-^")

	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDB.InitDB()
	dbMigrate(db)

	router := gin.Default()

	bookRpo := _bookRepo.NewPostgreSQLBookRepository(db)
	bookService := _bookService.NewBookService(bookRpo, 10)
	bookHandler := _bookHandler.NewBookHandler(bookService)

	routesInit := _routes.HandlerList{
		BookHandler: *bookHandler,
	}

	routesInit.RouteRegister(router)
}