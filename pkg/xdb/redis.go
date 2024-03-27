package xdb

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisDB struct {
	Address  string `json:"address"`
	DB       int    `json:"DB"`
	Password string `json:"password"`
}

func (r *RedisDB) RegisterRDBClient(c context.Context) (*redis.Client, error) {
	opt := redis.Options{
		Addr:     r.Address,
		DB:       r.DB,
		Password: r.Password,
	}
	cli := redis.NewClient(&opt)

	if err := cli.Ping(c).Err(); err != nil {
		return nil, err
	}

	return cli, nil
}
