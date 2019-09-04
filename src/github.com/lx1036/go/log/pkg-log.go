package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

/*
Gin Global Error Handler
 */

func logGin()  {
	gin.DisableConsoleColor()
	file, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(file)
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.String(200, "OK")
	})
	router.GET("ping", func(context *gin.Context) {
		context.String(200, "Pong")
	})
	router.Run(":8080")
}

func pkgLog()  {
	err := os.Chmod("./test.txt", 0777)
	if err != nil {
		fmt.Println("os chmod file failed")
		log.Fatal(err)
	}

	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file failed")
		log.Fatal(err)
	}

	err = file.Chmod(0777)
	if err != nil {
		fmt.Println("chmod failed")
		log.Fatal(err)
	}

	_, err = file.Write([]byte("abc"))
	if err != nil {
		fmt.Println("write bytes failed")
		log.Fatal(err, file.Fd())
	}
	/*
		write bytes failed
		2019/09/04 12:33:31 write ./test.txt: bad file descriptor 3
		exit status 1
	*/

	fmt.Println(file.Name())
}

func pkgErrors()  {
	err := errors.New("Error/Exception")
	if err != nil {
		fmt.Println(err)
	}
}


func main() {
	//pkgLog()

	//logGin()

	//pkgErrors()
}
