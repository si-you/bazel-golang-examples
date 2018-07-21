package main

import (
	"flag"
	"fmt"
	"github.com/mediocregopher/radix.v2/redis"
	"log"
)

var (
	redisAddr = flag.String("redis_addr", "localhost:6379", "Redis address.")
)

func main() {
	flag.Parse()

	c, err := redis.Dial("tcp", *redisAddr)
	if err != nil {
		log.Fatalf("Redis connection failed: %v", err)
	}

	resp := c.Cmd("SET", "hello", "world")
	if resp.Err != nil {
		log.Fatalf("Set message failed: %v", resp.Err)
	}
	fmt.Printf("SET completed: %v\n", resp)

	msg, err := c.Cmd("GET", "hello").Str()
	if err != nil {
		log.Fatalf("Get mesage failed: %v", err)
	}
	fmt.Printf("key: hello value: %v\n", msg)
}
