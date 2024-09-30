package utils

import (
	"context"
	"encoding/json"
	"errors"
	"task-inator3000/config"
	"time"

	"github.com/redis/go-redis/v9"
)

func AddToBlacklist(token string, c context.Context) error {
	key := blacklistKeyPrefix + token
	expiry := refreshTokenExp

	// defining the value for the key
	dataMap := map[string]interface{}{
		"blacklistedAt": time.Now().Unix(),
	}

	data, err := json.Marshal(dataMap)
	if err != nil {
		return errors.New("failed to marshal value")
	}

	err = config.RedisClient.Set(c, key, data, expiry).Err()
	if err != nil {
		return err
	}

	return nil
}

func IsBlacklisted(token string, c context.Context) (bool, error) {
	key := blacklistKeyPrefix + token

	_, err := config.RedisClient.Get(c, key).Result()
	if err != nil {
		// the following states that the key was not found
		if err == redis.Nil {
			return false, nil
		}

		// for other errors
		return false, err
	}

	return true, nil
}
