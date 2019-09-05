package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	Name string
	Age int
}

func main()  {
	app := gin.Default()
	routes := app.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": Person{
				Name: "lx1036",
				Age:  29,
			},
		})
	})

	routes.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"data": "world",
		})
	})
	
	err := app.Run(":8080")
	if err != nil {
		fmt.Printf("uncaught error: %v", err)
	}
}
