package radis

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v7"
	"gitlab.com/pragmaticreviews/golang-mux-api/entity"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewredisCache(host string, db int, exp time.Duration) {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
}
func (cache *redisCache) Set(key string, value *entity.Post) {
	client := cache.getClient()
	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(key,json,cache.expires*time.Second)
}
func (cache *redisCache) Get(key string) *entity.Post {
	client := cache.getClient()
	var,err := client.Get(key).Result()
	if err != nil{
		return nil
	}
	post := entity.Post{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil{
		panic(err)
	}
	return &post