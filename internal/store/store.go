package store

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Store interface {
	SaveShortenedURL(ctx context.Context, url string) (string, error)
	GetFullURL(ctx context.Context, code string) (string, error)
}

type store struct {
	rdb *redis.Client
}

func NewStore(rdb *redis.Client) Store {
	return &store{rdb: rdb}
}

func (s store) SaveShortenedURL(ctx context.Context, _url string) (string, error) {
	var code string
	for range 5 {
		code = generateCode()
		if err := s.rdb.HGet(ctx, "shortened_urls", code).Err(); err != nil {
			if err == redis.Nil {
				break
			}
			return "", fmt.Errorf("failed to save shortened url: %w", err)
		}
	}

	if err := s.rdb.HSet(ctx, "shortened_urls", code, _url).Err(); err != nil {
		return "", fmt.Errorf("failed to save shortened url: %w", err)
	}

	return code, nil
}

func (s store) GetFullURL(ctx context.Context, code string) (string, error) {
	fullUrl, err := s.rdb.HGet(ctx, "shortened_urls", code).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get full url: %w", err)
	}
	return fullUrl, nil
}
