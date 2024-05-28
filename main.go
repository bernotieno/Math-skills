package main

import (
	"fmt"
	"os"

	"Mathskill/processor"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Invalid input command, use: $ go run <your-program.go> <data.txt>")
		return
	}
	result, err := processor.Procession(args[1])
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Print(result)
}
