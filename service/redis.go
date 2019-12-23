package service

import (
    "context"
    "github.com/plexmediamanager/micro-redis/proto"
    serviceHandler "github.com/plexmediamanager/service"
    microClient "github.com/micro/go-micro/client"
    "time"
)

func GetRedisService(client microClient.Client) proto.RedisService {
    return proto.NewRedisService(serviceHandler.GetServiceName(serviceHandler.RedisServiceName), client)
}

func RedisHas(client microClient.Client, key string) (bool, error) {
    service := GetRedisService(client)
    lookupKey := &proto.GetRedisKey{}
    lookupKey.Key = key

    result, err := service.Has(context.TODO(), lookupKey)
    if err != nil {
        return false, err
    }
    return result.Response, err
}

func RedisSetValue(client microClient.Client, key string, value string, expiration time.Duration) (*proto.RedisEmpty, error) {
    service := GetRedisService(client)
    parameters := &proto.SetRedisObject{
        Key:                  key,
        Value:                value,
        Expiration:           int64(expiration),
    }
    return service.SetValue(context.TODO(), parameters)
}

func RedisGetValue(client microClient.Client, key string) (*proto.RedisResultBytes, error) {
    service := GetRedisService(client)
    parameters := &proto.GetRedisKey{
        Key:                  key,
    }
    return service.GetValue(context.TODO(), parameters)
}

func RedisGetValueAllowEmpty(client microClient.Client, key string) (*proto.RedisResultString, error) {
    service := GetRedisService(client)
    parameters := &proto.GetRedisKey{
        Key:                  key,
    }
    return service.GetValueAllowEmpty(context.TODO(), parameters)
}

func RedisGetValuesAsInt(client microClient.Client, key string) (*proto.RedisResultInt, error) {
    service := GetRedisService(client)
    parameters := &proto.GetRedisKey{
        Key:                  key,
    }
    return service.GetValueAsInt(context.TODO(), parameters)
}

func RedisGetValuesAsInt64(client microClient.Client, key string) (*proto.RedisResultInt64, error) {
    service := GetRedisService(client)
    parameters := &proto.GetRedisKey{
        Key:                  key,
    }
    return service.GetValueAsInt64(context.TODO(), parameters)
}

func RedisGetValuesAsUInt64(client microClient.Client, key string) (*proto.RedisResultUInt64, error) {
    service := GetRedisService(client)
    parameters := &proto.GetRedisKey{
        Key:                  key,
    }
    return service.GetValueAsUInt64(context.TODO(), parameters)
}

func RedisGetValuesAsFloat32(client microClient.Client, key string) (*proto.RedisResultFloat32, error) {
    service := GetRedisService(client)
    parameters := &proto.GetRedisKey{
        Key:                  key,
    }
    return service.GetValueAsFloat32(context.TODO(), parameters)
}

func RedisGetValuesAsFloat64(client microClient.Client, key string) (*proto.RedisResultFloat64, error) {
    service := GetRedisService(client)
    parameters := &proto.GetRedisKey{
        Key:                  key,
    }
    return service.GetValueAsFloat64(context.TODO(), parameters)
}

func RedisGetValuesAsTime(client microClient.Client, key string) (*proto.RedisResultUInt64, error) {
    service := GetRedisService(client)
    parameters := &proto.GetRedisKey{
        Key:                  key,
    }
    return service.GetValueAsTime(context.TODO(), parameters)
}

func RedisIncrement(client microClient.Client, key string) (*proto.RedisResultInt64, error) {
    service := GetRedisService(client)
    parameters := &proto.GetRedisKey{
        Key:                  key,
    }
    return service.Increment(context.TODO(), parameters)
}

func RedisIncrementBy(client microClient.Client, key string, value int64) (*proto.RedisResultInt64, error) {
    service := GetRedisService(client)
    parameters := &proto.IncrementDecrementRedis{
        Key:            key,
        Value:          value,
    }
    return service.IncrementBy(context.TODO(), parameters)
}

func RedisDecrement(client microClient.Client, key string) (*proto.RedisResultInt64, error) {
    service := GetRedisService(client)
    parameters := &proto.GetRedisKey{
        Key:                  key,
    }
    return service.Decrement(context.TODO(), parameters)
}

func RedisDecrementBy(client microClient.Client, key string, value int64) (*proto.RedisResultInt64, error) {
    service := GetRedisService(client)
    parameters := &proto.IncrementDecrementRedis{
        Key:            key,
        Value:          value,
    }
    return service.DecrementBy(context.TODO(), parameters)
}

func RedisRemove(client microClient.Client, keys ...string) (*proto.RedisResultInt64, error) {
    service := GetRedisService(client)
    parameters := &proto.GetRedisKeys{
        Keys:                  keys,
    }
    return service.Remove(context.TODO(), parameters)
}