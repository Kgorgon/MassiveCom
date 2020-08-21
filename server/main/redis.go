package main

import (
	"github.com/redis"
	"time"
)

var pool *redis.pool

func initPool(address string,maxIdle int,maxActive int,idleTimeout time.Duration){
	pool = &redis.Pool {}
		MaxIdle: maxIdle,//最大空闲连接数
		MaxActive: maxActive,//表示和数据库的最大连接数，0表示没有限制
		IdleTimeout:idleTimeout,//最大空闲时间
		Dial: func()(redis.Conn,error){//初始化连接代码，连接那个ip的redis
		return redis.Dial("tcp",address)
		},
	}
}