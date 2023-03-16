package dao

import (
	"fmt"
	"log"
	"mini-douyin/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DATABASE_USERNAME,
		config.DATABASE_PASSWORD,
		config.DATABASE_HOST,
		config.DATABASE_PORT,
		config.DATABASE_DBNAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln("数据库连接失败", err)
		panic(err)
	}
	Db = db

}
