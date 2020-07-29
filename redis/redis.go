package redis

import (
	"context"
	"go-jwt/conf"
	"net/url"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client
var ctx = context.Background()

func ConnectRedis() {
	u, err := url.Parse(conf.GetEnv("REDIS_URL"))
	if err != nil {
		panic(err.Error())
	}

	password, _ := u.User.Password()
	db, err := strconv.Atoi(u.Path[1:])
	if err != nil {
		panic("Redis url format error")
	}

	RDB = redis.NewClient(&redis.Options{
		Addr:     u.Host,
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}

func Get(key string) (string, error) {
	return RDB.Get(ctx, key).Result()
}

func Set(key string, value interface{}, dur time.Duration) error {
	return RDB.Set(ctx, key, value, dur).Err()
}

func Exists(key string) bool {
	ok, _ := RDB.Exists(ctx, key).Result()
	return ok == 1
}
