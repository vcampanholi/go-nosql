package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"go_nosql/app/routes"
)

func main() {
	router := gin.Default()

	router.POST("/routes", routes.Create)

	err := router.Run("localhost: 3000")
	if err != nil {
		panic(errors.New("error start app"))
	}
}
