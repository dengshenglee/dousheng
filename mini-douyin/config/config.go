package config

import "time"

const GO_STARTER_TIME = "2006-01-02 15:04:05"

// DATABASE
const (
	DATABASE_USERNAME = "user"
	DATABASE_PASSWORD = "gorm"
	DATABASE_HOST     = "localhost"
	DATABASE_PORT     = 9910
	DATABASE_DBNAME   = "dousheng"
)

// Redis
const (
	ProdRedisAddr = "localhost:16379"
	ProRedisPwd   = "dousheng"
)

// videos
const (
	VIDEO_NUM_PER_REFRESH     = 6
	VIDEO_INIT_NUM_PER_AUTHOR = 10
	VIDEO_COVER_IMAGE         = "XXX"
)

// OSS config
// const (
// 	OSS_ACCESS_KEY_ID     = "OSS_ACCESS_KEY_ID"
// 	OSS_ACCESS_KEY_SECRET = "OSS_ACCESS_KEY_SECRET"
// 	OSS_BUCKET_NAME       = "OSS_BUCKET_NAME"
// 	OSS_ENDPOINT          = "OSS_ENDPOINT"
// 	CUSTOM_DOMAIN         = "CUSTOM_DOMAIN"
// 	OSS_VIDEO_DIR         = "OSS_VIDEO_DIR"
// 	PLAY_URL_PREFIX       = CUSTOM_DOMAIN + OSS_VIDEO_DIR
// 	COVER_URL_SUFFIX      = "?x-oss-process=video/snapshot,t_2000,m_fast"
// )

const (
	OSS_ENDPOINT          = "localhost:9000"
	OSS_ACCESS_KEY_ID     = "admin"
	OSS_SECRET_ACCESS_KEY = "admin123"
	OSS_useSSL            = false
	OSS_BUCKET_NAME       = "dousheng"
	OSS_LOCATION          = "us-east-1"
	//TODO: 获取用户头像地址
	CUSTOM_DOMAIN       = "xxx"
	OSS_USER_AVATAR_DIR = "xxx"
	BG_IMAGE            = "url://BackgroundImage.png"
	SIGNATURE           = "god bless you"
)

// jwt key and expire time
const (
	JWT_KEY         = "dousheng"
	JWT_EXPIRE_TIME = 24 * time.Hour
)
