package db

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

var DB *redis.Client