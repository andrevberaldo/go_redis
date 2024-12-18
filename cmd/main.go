package main

import (
	"fmt"

	"github.com/andrevberaldo/go_redis/usecase"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Unable to load environments variables\n")
	}
}

func main() {
	userId := 2
	getUserPostsUseCase := usecase.NewGetUserPostsUseCase(userId)

	userPosts, err := getUserPostsUseCase.Execute()

	if err != nil {
		fmt.Println(err.Error())
	}

	//  index, post
	fmt.Printf("Posts Length %d\n", len(userPosts))
	for _, post := range userPosts {
		parsedPost, err := post.ToJSON()

		if err != nil {
			fmt.Println("Unable to parse []byte to JSON")
		}

		fmt.Println(string(parsedPost))
	}
}
