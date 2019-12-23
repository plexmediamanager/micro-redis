package resolver

import (
    "context"
    "github.com/plexmediamanager/micro-redis/proto"
    "github.com/plexmediamanager/micro-redis/redis"
    "strconv"
    "time"
)

type RedisService struct {
    Redis           *redis.RedisClient
}

func (service RedisService) errorToProto(err error) *proto.RedisEmpty {
    return &proto.RedisEmpty{}
}

func (service RedisService) SetValue(_ context.Context, parameters *proto.SetRedisObject, response *proto.RedisEmpty) error {
    err := service.Redis.SetValue(parameters.Key, parameters.Value, time.Duration(parameters.Expiration))
    if err != nil {
        return err
    }
    return nil
}

func (service RedisService) Get (_ context.Context, parameters *proto.GetRedisKey, response *proto.RedisResultBytes) error {
    result, err := service.Redis.Get(parameters.Key)
    if err != nil {
        return err
    }
    response.Response = result
    return nil
}

func (service RedisService) GetValue (_ context.Context, parameters *proto.GetRedisKey, response *proto.RedisResultBytes) error {
    result, err := service.Redis.GetValue(parameters.Key)
    if err != nil {
        return err
    }
    response.Response = result
    return nil
}

func (service RedisService) GetValueAllowEmpty (_ context.Context, parameters *proto.GetRedisKey, response *proto.RedisResultString) error {
    response.Response = service.Redis.GetValueAllowEmpty(parameters.Key)
    return nil
}

func (service RedisService) GetValueAsInt (_ context.Context, parameters *proto.GetRedisKey, response *proto.RedisResultInt) error {
    result, err := service.Redis.GetValueAsInt(parameters.Key)
    if err != nil {
        return err
    }
    response.Response = int32(result)
    return nil
}

func (service RedisService) GetValueAsInt64 (_ context.Context, parameters *proto.GetRedisKey, response *proto.RedisResultInt64) error {
    result, err := service.Redis.GetValueAsInt64(parameters.Key)
    if err != nil {
        return err
    }
    response.Response = result
    return nil
}

func (service RedisService) GetValueAsUInt64 (_ context.Context, parameters *proto.GetRedisKey, response *proto.RedisResultUInt64) error {
    result, err := service.Redis.GetValueAsUInt64(parameters.Key)
    if err != nil {
        return err
    }
    response.Response = result
    return nil
}

func (service RedisService) GetValueAsFloat32 (_ context.Context, parameters *proto.GetRedisKey, response *proto.RedisResultFloat32) error {
    result, err := service.Redis.GetValueAsFloat32(parameters.Key)
    if err != nil {
        return err
    }
    response.Response = result
    return nil
}

func (service RedisService) GetValueAsFloat64 (_ context.Context, parameters *proto.GetRedisKey, response *proto.RedisResultFloat64) error {
    result, err := service.Redis.GetValueAsFloat64(parameters.Key)
    if err != nil {
        return err
    }
    response.Response = result
    return nil
}

func (service RedisService) GetValueAsTime (_ context.Context, parameters *proto.GetRedisKey, response *proto.RedisResultUInt64) error {
    result, err := service.Redis.GetValueAsTime(parameters.Key)
    if err != nil {
        return err
    }
    value, err := strconv.ParseUint(result.String(), 10, 64)
    if err != nil {
        return err
    }
    response.Response = value
    return nil
}

func (service RedisService) Increment (_ context.Context, parameters *proto.GetRedisKey, response *proto.RedisResultInt64) error {
    result, err := service.Redis.Increment(parameters.Key)
    if err != nil {
        return err
    }
    response.Response = result
    return nil
}

func (service RedisService) IncrementBy (_ context.Context, parameters *proto.IncrementDecrementRedis, response *proto.RedisResultInt64) error {
    result, err := service.Redis.IncrementBy(parameters.Key, parameters.Value)
    if err != nil {
        return err
    }
    response.Response = result
    return nil
}

func (service RedisService) Decrement (_ context.Context, parameters *proto.GetRedisKey, response *proto.RedisResultInt64) error {
    result, err := service.Redis.Decrement(parameters.Key)
    if err != nil {
        return err
    }
    response.Response = result
    return nil
}

func (service RedisService) DecrementBy (_ context.Context, parameters *proto.IncrementDecrementRedis, response *proto.RedisResultInt64) error {
    result, err := service.Redis.IncrementBy(parameters.Key, parameters.Value)
    if err != nil {
        return err
    }
    response.Response = result
    return nil
}

func (service RedisService) Has (_ context.Context, parameters *proto.GetRedisKey, response *proto.RedisResultBoolean) error {
    response.Response = service.Redis.Has(parameters.Key)
    return nil
}

func (service RedisService) Remove (_ context.Context, parameters *proto.GetRedisKeys, response *proto.RedisResultInt64) error {
    result, err := service.Redis.Remove(parameters.Keys...)
    if err != nil {
        return err
    }
    response.Response = result
    return nil
}