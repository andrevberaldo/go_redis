package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Unable to load environments variables")
	}
}

func main() {
	fmt.Printf("Hello ::GO with Redis")
}
