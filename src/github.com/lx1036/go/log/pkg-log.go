package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
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
