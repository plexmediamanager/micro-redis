package main

import (
    "github.com/micro/go-micro/client"
    "github.com/plexmediamanager/micro-redis/proto"
    "github.com/plexmediamanager/micro-redis/redis"
    "github.com/plexmediamanager/micro-redis/resolver"
    "github.com/plexmediamanager/service"
    "github.com/plexmediamanager/service/log"
    "time"
)

func main() {
    application := service.CreateApplication()
    redisClient, err := redis.Initialize().Connect()
    if err != nil {
        log.Panic(err)
    }

    err = application.InitializeConfiguration()
    if err != nil {
        log.Panic(err)
    }

    err = application.InitializeMicroService()
    if err != nil {
        log.Panic(err)
    }

    err = application.Service().Client().Init(
        client.PoolSize(10),
        client.Retries(30),
        client.RequestTimeout(1 * time.Second),
    )
    if err != nil {
        log.Panic(err)
    }

    err = proto.RegisterRedisServiceHandler(application.Service().Server(), resolver.RedisService{ Redis: redisClient })
    if err != nil {
        log.Panic(err)
    }

    go application.StartMicroService()

    service.WaitForOSSignal(1)

    err = redisClient.Disconnect()
    if err != nil {
        log.Printf("Connection to Redis was closed with the following error: %v", err)
    }
}
