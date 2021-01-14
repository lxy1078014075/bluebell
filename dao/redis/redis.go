package redis

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			viper.GetString("redis.host"),
			viper.GetInt("redis.port"),
		),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
	})
	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}

func SetKey(key, value string) error {
	return rdb.Set(key, value, 0).Err()
}

func GetKey(key string) (string, error) {
	value, err := rdb.Get(key).Result()
	if err == redis.Nil {
		return "", errors.New("key not exist")
	} else if err != nil {
		return "", err
	}
	return value, err
}
