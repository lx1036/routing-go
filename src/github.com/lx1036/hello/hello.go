package main

import (
	"fmt"
	"github.com/lx1036/stringutil"
	)

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	fmt.Printf(stringutil.Reverse("Hello, world."))
}
