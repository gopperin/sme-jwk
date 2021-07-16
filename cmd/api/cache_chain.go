package api

import (
	"time"

	"github.com/dgraph-io/ristretto"
	"github.com/eko/gocache/cache"
	"github.com/eko/gocache/store"
	"github.com/go-redis/redis/v7"
	"github.com/sirupsen/logrus"
)

// InitChainCache InitChainCache
func InitChainCache() (*cache.ChainCache, error) {

	// Initialize Ristretto cache and Redis client
	ristrettoCache, err := ristretto.NewCache(&ristretto.Config{NumCounters: 1000, MaxCost: 100, BufferItems: 64})
	if err != nil {
		logrus.Println(err.Error())
		panic(err)
	}

	redisClient := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})

	// Initialize stores
	ristrettoStore := store.NewRistretto(ristrettoCache, nil)
	redisStore := store.NewRedis(redisClient, &store.Options{Expiration: 5 * time.Second})

	// Initialize chained cache
	_cacheManager := cache.NewChain(
		cache.New(ristrettoStore),
		cache.New(redisStore),
	)

	return _cacheManager, nil
}
