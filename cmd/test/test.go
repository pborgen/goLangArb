package main

import (
	"fmt"

	"github.com/paulborgen/goLangArb/internal/cache"
)

func main() {
	fmt.Println("Hello123")

	cache := cache.CreateCache[string, int](3)
	cache.Put("1", 1)
	cache.Put("2", 2)
	cache.Put("3", 3)
	cache.Put("4", 4)
	cache.Put("5", 5)

	fmt.Println(cache.Get("4"))
}
