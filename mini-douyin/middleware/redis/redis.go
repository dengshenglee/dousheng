package redis

import (
	"context"
	"mini-douyin/config"

	"github.com/redis/go-redis/v9"
)

var (
	Ctx      = context.Background()
	NilError = redis.Nil
	RdbTest  *redis.Client
)

// UserFollowings 根据用户id找到他关注的人
var UserFollowings *redis.Client

// UserFollowers 根据用户id找到他的粉丝
var UserFollowers *redis.Client

// UserFriends 根据用户id找到他的好友
var UserFriends *redis.Client

// RdbVCid 存储video与comment的关系
var RdbVCid *redis.Client

// RdbCVid 根据commentId找videoId
var RdbCVid *redis.Client

// RdbCIdComment 根据commentId 找comment
var RdbCIdComment *redis.Client

// RdbUVid 根据userId找到他点赞过的videoId
var RdbUVid *redis.Client

// RdbVUid 根据videoId找到点赞过它的userId
var RdbVUid *redis.Client

// InitRedis 初始化 Redis 连接，redis 默认 16 个 DB
func InitRedis() {
	RdbTest = redis.NewClient(&redis.Options{
		Addr:     config.ProdRedisAddr,
		Password: config.ProRedisPwd,
		DB:       0,
	})
	RdbVCid = redis.NewClient(&redis.Options{
		Addr:     config.ProdRedisAddr,
		Password: config.ProRedisPwd,
		DB:       1,
	})
	RdbCVid = redis.NewClient(&redis.Options{
		Addr:     config.ProdRedisAddr,
		Password: config.ProRedisPwd,
		DB:       2,
	})
	RdbCIdComment = redis.NewClient(&redis.Options{
		Addr:     config.ProdRedisAddr,
		Password: config.ProRedisPwd,
		DB:       3,
	})
	RdbUVid = redis.NewClient(&redis.Options{
		Addr:     config.ProdRedisAddr,
		Password: config.ProRedisPwd,
		DB:       4,
	})
	RdbVUid = redis.NewClient(&redis.Options{
		Addr:     config.ProdRedisAddr,
		Password: config.ProRedisPwd,
		DB:       5,
	})
	UserFollowings = redis.NewClient(&redis.Options{
		Addr:     config.ProdRedisAddr,
		Password: config.ProRedisPwd,
		DB:       11,
	})
	UserFollowers = redis.NewClient(&redis.Options{
		Addr:     config.ProdRedisAddr,
		Password: config.ProRedisPwd,
		DB:       12,
	})
	UserFriends = redis.NewClient(&redis.Options{
		Addr:     config.ProdRedisAddr,
		Password: config.ProRedisPwd,
		DB:       13,
	})
}
