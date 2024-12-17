package main

import (
	"fmt"

	"github.com/andrevberaldo/go_redis/usecase"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Unable to load environments variables")
	}
}

func main() {
	getUserPostsUseCase := usecase.NewGetUserPostsUseCase(2)

	userPosts, err := getUserPostsUseCase.Execute()

	if err != nil {
		fmt.Println(err.Error())
	}

	//  index, post
	for _, post := range userPosts {
		json, err := post.ToJSON()

		if err != nil {
			fmt.Println("Unable to parse []byte to JSON")
		}

		fmt.Println(string(json))
	}
}
