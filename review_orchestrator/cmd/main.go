package main

import (
	"fmt"

	"github.com/ullas-shivakumar/llmCodeReviewer/internal/cache"
)

func main() {
	fmt.Println("Initializng Redis Cache:")
	cache.ConnectRedisCache()
}
