package main

import (
	"os"
	"fmt"
)

// import "log" package

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("err:", err)
		// log.Println()
		// log.Fatalln(err)
		// panic(err)
	}
}