package redis

import (
	"encoding/json"
	"time"
	"user-backend/conf"
	"user-backend/models"

	"github.com/go-redis/redis"
)

type Client interface {
	Set(key string, values interface{}) error
	Get(key string) (string, error)
	Del(key string) error
	UpdateApiLimiter(key string, values interface{}) error
}

type client struct {
	client *redis.Client
}

func NewClient(conf *conf.Config) Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     conf.Redis.Addr,
		Password: conf.Redis.Password,
		DB:       1,
	})
	if redisClient == nil {
		panic("failed to connect to redis")
	}

	return &client{
		client: redisClient,
	}
}

func (r *client) Set(key string, values interface{}) error {
	err := r.client.Set(key, values, time.Minute*5).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *client) Get(key string) (string, error) {
	res := r.client.Get(key)
	if res.Err() == redis.Nil {
		return "", nil

	} else if res.Val() != "" {
		return res.Val(), nil
	}
	return "", res.Err()
}

func (r *client) Del(key string) error {
	err := r.client.Del(key).Err()
	if err != nil {
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *client) SetApiLimiter(key string, values interface{}) error {
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}
	err = r.client.Set(key, data, time.Minute*60).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *client) GetSetApiLimiter(key string, values interface{}) error {
	data, err := json.Marshal(values)
	if err != nil {
		return err
	}
	ttl := r.client.TTL(key).Val().Seconds()
	err = r.client.Set(key, data, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *client) UpdateApiLimiter(key string, values interface{}) error {
	val, err := r.Get(key)
	if err != nil {
		return err
	}
	if len(val) == 0 {
		err = r.SetApiLimiter(key, values)
		if err != nil {
			return err
		}
	} else {
		var otpValue models.APILimiterDto
		err = json.Unmarshal([]byte(val), &otpValue)
		if err != nil {

			return err
		}
		otpValue.Tries = otpValue.Tries + 1
		err = r.GetSetApiLimiter(key, otpValue)
		if err != nil {
			return err
		}
	}
	return nil
}
