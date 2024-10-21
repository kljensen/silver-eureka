package main

import (
	"fmt"
	"os"
	"strconv"
)

func printHelloWorld(count int) {
	for i := 0; i < count; i++ {
		fmt.Println("hello world")
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide the number of times to print 'hello world'")
		return
	}

	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid number:", os.Args[1])
		return
	}
	printHelloWorld(count)
}
