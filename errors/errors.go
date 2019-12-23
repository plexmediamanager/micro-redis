package errors

import "github.com/plexmediamanager/service/errors"

const (
    ServiceID       errors.Service      =   1
)

var (
    InvalidConfigurationError = errors.Error {
        Code:       errors.Code {
            Service: ServiceID,
            ErrorType: errors.TypeService,
            ErrorNumber: 1,
        },
        Message:    "Redis service is not configured properly, please check your configuration",
    }
    RedisConnectionError = errors.Error {
        Code:       errors.Code {
            Service: ServiceID,
            ErrorType: errors.TypeWrapper,
            ErrorNumber: 1,
        },
        Message:    "Redis service is not able to reach specified Redis instance(s)",
    }
    RedisSetValueError = errors.Error {
        Code:       errors.Code {
            Service: ServiceID,
            ErrorType: errors.TypeWrapper,
            ErrorNumber: 2,
        },
        Message:    "Unable to set value at key: %s",
    }
    RedisGetValueError = errors.Error {
        Code:       errors.Code {
            Service: ServiceID,
            ErrorType: errors.TypeWrapper,
            ErrorNumber: 3,
        },
        Message:    "Unable to get value at key: %s",
    }
    RedisIncrementError = errors.Error {
        Code:       errors.Code {
            Service: ServiceID,
            ErrorType: errors.TypeWrapper,
            ErrorNumber: 4,
        },
        Message:    "Unable to increment value at key: %s",
    }
    RedisIncrementByError = errors.Error {
        Code:       errors.Code {
            Service: ServiceID,
            ErrorType: errors.TypeWrapper,
            ErrorNumber: 5,
        },
        Message:    "Unable to increment value at key `%s` by `%d`",
    }
    RedisDecrementError = errors.Error {
        Code:       errors.Code {
            Service: ServiceID,
            ErrorType: errors.TypeWrapper,
            ErrorNumber: 6,
        },
        Message:    "Unable to decrement value at key: %s",
    }
    RedisDecrementByError = errors.Error {
        Code:       errors.Code {
            Service: ServiceID,
            ErrorType: errors.TypeWrapper,
            ErrorNumber: 7,
        },
        Message:    "Unable to decrement value at key `%s` by `%d`",
    }
    RedisRemoveKeysError = errors.Error {
        Code:       errors.Code {
            Service: ServiceID,
            ErrorType: errors.TypeWrapper,
            ErrorNumber: 8,
        },
        Message:    "Unable to delete these keys: `%v`",
    }
)
