package redis

import (
	"github.com/gomodule/redigo/redis"
)

// Connection
func Connection() redis.Conn {
	const Addr = "127.0.0.1:6379"

	c, err := redis.Dial("tcp", Addr)
	if err != nil {
		panic(err)
	}
	return c
}

// Set データの登録(Redis: SET key value)
func Set(key, value string, c redis.Conn) string {
	res, err := redis.String(c.Do("SET", key, value))
	if err != nil {
		panic(err)
	}
	return res
}

// Get データの取得(Redis: GET key)
func Get(key string, c redis.Conn) string {
	res, err := redis.String(c.Do("GET", key))
	if err != nil {
		panic(err)
	}
	return res
}
