package main

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
	"time"
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

/*
SENTRY_DSN, SENTRY_RELEASE, SENTRY_ENVIRONMENT
 */
func sentryDemo()  {
	viper.AddConfigPath(".")
	viper.SetConfigFile("env.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Read config file failed:%v\n", err))
	}
	sentryDsn := viper.Get("sentry_dsn")
	//fmt.Println(sentryDsn, viper.AllKeys())

	err = sentry.Init(sentry.ClientOptions{
		Dsn:              cast.ToString(sentryDsn),
		Debug:            true,
		AttachStacktrace: true,
		SampleRate:       0,
		IgnoreErrors:     nil,
		BeforeSend:       nil,
		BeforeBreadcrumb: nil,
		Integrations:     nil,
		DebugWriter:      nil,
		Transport:        nil,
		ServerName:       "",
		Release:          "",
		Dist:             "",
		Environment:      "",
		MaxBreadcrumbs:   0,
		HTTPTransport:    nil,
		HTTPProxy:        "",
		HTTPSProxy:       "",
		CaCerts:          nil,
	})
	if err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	_, err = os.Open("./sentry2.txt")
	if err != nil {
		fmt.Printf("Open file failed: %v\n", err)
		/*
		https://sentry.io/organizations/leftcapital/issues/?project=1550933&query=is%3Aunresolved&statsPeriod=14d
		 */
		sentry.CaptureException(err)
		sentry.Flush(time.Second * 5)
	}
}

func main() {
	//pkgLog()

	//logGin()

	//pkgErrors()

	sentryDemo()
}
