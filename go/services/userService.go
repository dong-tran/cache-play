package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/dong-tran/cache-play/model"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type userService struct {
	key   string
	db    *gorm.DB
	redis *redis.Client
}

type UserService interface {
	GetUsers(partition int64) ([]model.Users, error)
}

func NewUserService(redis *redis.Client, db *gorm.DB) UserService {
	return &userService{
		key:   "api/users",
		redis: redis,
		db:    db,
	}
}

func (s *userService) GetUsers(partition int64) ([]model.Users, error) {
	key := fmt.Sprintf("%s/%d", s.key, partition)
	ctx := context.Background()
	result := make([]model.Users, 0)
	val, err := s.redis.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return result, err
	}
	if len(val) > 0 {
		err := json.Unmarshal([]byte(val), &result)
		if err != nil {
			return result, err
		}
		return result, nil
	} else {
		from := (partition-1)*100 + 1
		to := (partition * 100)
		s.db.Where("id BETWEEN ? AND ?", from, to).Find(&result)
		str, _ := json.Marshal(result)
		s.redis.Set(ctx, key, string(str), 60*time.Second)
		return result, nil
	}
}
