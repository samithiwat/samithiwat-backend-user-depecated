package repository

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/samithiwat/samithiwat-backend-user/src/proto"
	"gorm.io/gorm"
	"math"
	"time"
)

func Pagination(value interface{}, meta *proto.PaginationMetadata) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		var totalItems int64
		db.Model(value).Count(&totalItems)

		meta.TotalItem = totalItems
		totalPages := math.Ceil(float64(totalItems) / float64(meta.ItemsPerPage))
		meta.TotalPage = int64(totalPages)

		if meta.CurrentPage < 1 {
			meta.CurrentPage = 1
		}

		switch {
		case meta.ItemsPerPage > 100:
			meta.ItemsPerPage = 100
		case meta.ItemsPerPage < 10:
			meta.ItemsPerPage = 10
		}

		offset := (meta.CurrentPage - 1) * meta.ItemsPerPage
		return db.Offset(int(offset)).Limit(int(meta.ItemsPerPage))
	}
}

func SaveCache(c *redis.Client, key string, value interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	v, err := json.Marshal(value)
	if err != nil {
		return
	}

	return c.Set(ctx, key, v, 0).Err()
}

func GetCache(c *redis.Client, key string, value interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	v, err := c.Get(ctx, key).Result()
	if err != nil {
		return
	}

	return json.Unmarshal([]byte(v), value)
}
