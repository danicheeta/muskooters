package redis

import (
	"github.com/go-redis/redis"
	"muskooters/services/config"
	"octapus/services/assert"
	"octapus/services/initializer"
)

var Client *redis.Client

type manager struct{}

func (manager) Initialize() func() {
	redisURL := config.MustString("REDIS_URL")

	Client = redis.NewClient(&redis.Options{
		Addr: redisURL,
		DB:   0,
	})
	_, err := Client.Ping().Result()
	assert.Nil(err)

	return nil
}

func init() {
	initializer.Register(manager{})
}
