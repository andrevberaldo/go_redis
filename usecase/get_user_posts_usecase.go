package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/andrevberaldo/go_redis/model"
	"github.com/andrevberaldo/go_redis/services"
)

type GetUserPostsUseCase struct {
	UserId   uint32
	PostsURL string
}

func NewGetUserPostsUseCase(id uint32) GetUserPostsUseCase {
	return GetUserPostsUseCase{
		UserId:   id,
		PostsURL: fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", id),
	}
}

func (u *GetUserPostsUseCase) Execute() ([]model.Post, error) {
	var posts []model.Post

	// check first if the cache has the response

	// if so, return response from cache

	// if not so, fetch data in API, set it as cache and return the response
	body, err := services.FetchData(u.PostsURL)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &posts)

	return posts, nil
}
