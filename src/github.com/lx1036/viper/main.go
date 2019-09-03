package main

import (
	"fmt"
	config "github.com/lx1036/viper/test"
	"github.com/spf13/viper"
)

func main()  {
	if err := config.Init(); err != nil {
		panic(err)
	}

	name := viper.GetString("name")
	fmt.Println("Viper name:", name)
}
