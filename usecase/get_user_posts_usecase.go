package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/andrevberaldo/go_redis/model"
	"github.com/andrevberaldo/go_redis/services"
	"github.com/redis/go-redis/v9"
)

type GetUserPostsUseCase struct {
	UserId   int
	PostsURL string
	Cache    redis.Client
}

var ctx = context.Background()

func NewGetUserPostsUseCase(id int) GetUserPostsUseCase {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return GetUserPostsUseCase{
		UserId:   id,
		PostsURL: fmt.Sprintf("https://jsonplaceholder.typicode.com/posts?userId=%d", id),
		Cache:    *redis,
	}
}

func (u *GetUserPostsUseCase) Execute() ([]model.Post, error) {
	var posts []model.Post

	// check first if the cache has the response
	cacheValue, err := u.Cache.Get(ctx, u.PostsURL).Result()

	// if not so, fetch data in API, set it as cache and return the response
	if err == redis.Nil {
		fmt.Println("Fetching data from API")

		apiValue, err := services.FetchData(u.PostsURL)

		if err != nil {
			fmt.Printf("Failed to get data from API: %p\n", err)
			return nil, err
		}

		js := string(apiValue)

		// parse api response to struct
		if err := json.Unmarshal([]byte(apiValue), &posts); err != nil {
			fmt.Printf("Failed to unmarshal API data: %p\n", err)
		}

		// set to cache
		u.Cache.Set(ctx, u.PostsURL, js, 10*time.Minute)

		return posts, nil

	} else if err != nil {
		fmt.Printf("Failed to get data from Cache: %p\n", err)
		return nil, err
	}

	fmt.Println("Returning Value from Cache")

	if err := json.Unmarshal([]byte(cacheValue), &posts); err != nil {
		fmt.Printf("Failed to unmarshal cache data: %p\n", err)
	}

	return posts, nil
}
