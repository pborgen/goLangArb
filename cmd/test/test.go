package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello123")

	args := os.Args
	fmt.Printf("Type of Args = %T\n", args)
	fmt.Println(args[0], args[1])
}
