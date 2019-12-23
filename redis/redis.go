package redis

import (
    "github.com/plexmediamanager/micro-redis/errors"
    format "fmt"
    "github.com/plexmediamanager/service/helpers"
    "github.com/go-redis/redis/v7"
    "strings"
    "time"
)

type RedisClient struct {
    host            string
    port            int
    endpoints       []string
    password        string
    database        int
    poolSize        int
    prefix          string

    connection      redis.UniversalClient
}

// Initialize Redis RedisClient
func Initialize() *RedisClient {
    client := &RedisClient{
        password:   helpers.GetEnvironmentVariableAsString("REDIS_PASSWORD", ""),
        database:   helpers.GetEnvironmentVariableAsInteger("REDIS_DATABASE", 0),
        poolSize:   helpers.GetEnvironmentVariableAsInteger("REDIS_POOL_SIZE", 100),
        prefix:     helpers.GetEnvironmentVariableAsString("REDIS_PREFIX", ""),
        connection: nil,
    }

    multipleEndpoints := helpers.GetEnvironmentVariableAsString("REDIS_ENDPOINTS", "")
    if len(multipleEndpoints) >= 9 {
        var endpoints []string
        for _, value := range strings.Split(multipleEndpoints, ",") {
            endpoints = append(endpoints, strings.TrimSpace(value))
        }
        client.endpoints = endpoints
    } else {
        client.host = helpers.GetEnvironmentVariableAsString("REDIS_HOST", "127.0.0.1")
        client.port = helpers.GetEnvironmentVariableAsInteger("REDIS_PORT", 6379)
    }

    return client
}

// Connect to Redis
func (client *RedisClient) Connect() (*RedisClient, error) {
    if client.host == "" && len(client.endpoints) == 0 {
        return nil, errors.InvalidConfigurationError.ToError(nil)
    }
    client.connection = redis.NewUniversalClient(&redis.UniversalOptions {
        Addrs:      client.buildRedisEndpointsConfiguration(),
        Password:   client.password,
        DB:         client.database,
        PoolSize:   client.poolSize,
    })

    err := client.connection.Ping().Err()
    if err != nil {
        return nil, errors.RedisConnectionError.ToError(err)
    }
    return client, nil
}

// Disconnect from Redis
func (client *RedisClient) Disconnect() error {
    if client.connection != nil {
        return client.connection.Close()
    }
    return nil
}

// Write value to Redis
func (client *RedisClient) SetValue(key string, value string, expiration time.Duration) error {
    err := client.connection.Set(client.buildRedisKey(key), value, expiration).Err()
    if err != nil {
        return errors.RedisSetValueError.ToErrorWithArguments(err, key)
    }
    return nil
}

func (client *RedisClient) Get(key string) ([]byte, error) {
    result, err := client.getValue(key).Bytes()
    return result, client.formatReadError(err, key)
}

// Get value from Redis (as a plain string)
func (client *RedisClient) GetValue(key string) ([]byte, error) {
    result, err := client.getValue(key).Bytes()
    return result, client.formatReadError(err, key)
}

// Get value from Redis (as a plain string AND ignore error)
func (client *RedisClient) GetValueAllowEmpty(key string) string {
    return client.getValue(key).String()
}

// Get value from Redis as Int
func (client *RedisClient) GetValueAsInt(key string) (int, error) {
    result, err := client.getValue(key).Int()
    return result, client.formatReadError(err, key)
}

// Get value from Redis as Int64
func (client *RedisClient) GetValueAsInt64(key string) (int64, error) {
    result, err := client.getValue(key).Int64()
    return result, client.formatReadError(err, key)
}

// Get value from Redis as UInt64
func (client *RedisClient) GetValueAsUInt64(key string) (uint64, error) {
    result, err := client.getValue(key).Uint64()
    return result, client.formatReadError(err, key)
}

// Get value from Redis as Float32
func (client *RedisClient) GetValueAsFloat32(key string) (float32, error) {
    result, err := client.getValue(key).Float32()
    return result, client.formatReadError(err, key)
}

// Get value from Redis as Float64
func (client *RedisClient) GetValueAsFloat64(key string) (float64, error) {
    result, err := client.getValue(key).Float64()
    return result, client.formatReadError(err, key)
}

// Get value from Redis as Time
func (client *RedisClient) GetValueAsTime(key string) (time.Time, error) {
    result, err := client.getValue(key).Time()
    return result, client.formatReadError(err, key)
}

// Increment provided key value
func (client *RedisClient) Increment(key string) (int64, error) {
    newNumber, err := client.connection.Incr(client.buildRedisKey(key)).Result()
    if err != nil {
        err = errors.RedisIncrementError.ToErrorWithArguments(err, key)
    }
    return newNumber, err
}

// Increment by provided key value
func (client *RedisClient) IncrementBy(key string, value int64) (int64, error) {
    newNumber, err := client.connection.IncrBy(client.buildRedisKey(key), value).Result()
    if err != nil {
        err = errors.RedisIncrementByError.ToErrorWithArguments(err, key, value)
    }
    return newNumber, err
}

// Decrement provided key value
func (client *RedisClient) Decrement(key string) (int64, error) {
    newNumber, err := client.connection.Decr(client.buildRedisKey(key)).Result()
    if err != nil {
        err = errors.RedisDecrementError.ToErrorWithArguments(err, key)
    }
    return newNumber, err
}

// Decrement provided key value
func (client *RedisClient) DecrementBy(key string, value int64) (int64, error) {
    newNumber, err := client.connection.DecrBy(client.buildRedisKey(key), value).Result()
    if err != nil {
        err = errors.RedisDecrementByError.ToErrorWithArguments(err, key, value)
    }
    return newNumber, err
}

// Check if Redis has provided key
func (client *RedisClient) Has(key string) bool {
    _, err := client.GetValue(key)
    return err == nil
}

// Remove key(s) from Redis
func (client *RedisClient) Remove(keys ...string) (int64, error) {
    var keysList []string
    for _, key := range keys {
        keysList = append(keysList, client.buildRedisKey(key))
    }
    removedKeysCount, err := client.connection.Del(keys...).Result()
    if err != nil {
        err = errors.RedisRemoveKeysError.ToErrorWithArguments(err, keys)
    }
    return removedKeysCount, err
}

// Create Redis connection address string
func (client *RedisClient) createRedisAddress() string {
    return format.Sprintf("%s:%d", client.host, client.port)
}

// Build endpoints configuration for Redis
func (client *RedisClient) buildRedisEndpointsConfiguration() []string {
    var endpoints []string
    if len(client.endpoints) > 0 {
        endpoints = client.endpoints
    } else {
        endpoints = append(endpoints, client.createRedisAddress())
    }
    return endpoints
}

// Get value without casting it to type
func (client *RedisClient) getValue(key string) *redis.StringCmd {
    return client.connection.Get(client.buildRedisKey(key))
}

// Build Redis key
func (client *RedisClient) buildRedisKey(key string) string {
    if len(client.prefix) == 0 {
        return key
    }
    return format.Sprintf("%s::%s", client.prefix, key)
}

// Format redis read error
func (client *RedisClient) formatReadError(err error, key string) error {
    if err != nil {
        err = errors.RedisGetValueError.ToErrorWithArguments(err, key)
    }
    return err
}