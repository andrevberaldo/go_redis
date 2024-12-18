## Go Sandbox to practice integration between Go and Redis

## How to Run
- Run `docker compose up`. It will setup localhost redis (on port 6379) and redis Insights (on port 5540)
- Run `go run ./cmd/main.go`

## Explanation
If the user's posts was never fetched before, the code will get them from the api, set as cache and return the value. So the next time the same URL is fetched and the cache is not expired, the code will return the value from cache istead of call the API again. This will increase the performance. 
For AWS cloud, we can use **ElastiCache** to achieve the same result with some code changes.