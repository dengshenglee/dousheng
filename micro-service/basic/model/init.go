package model

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const MySQLDefaultDSN = "user:gorm@tcp(localhost:9910)/dousheng?charset=utf8&parseTime=True&loc=Local"

var (
	DB   *gorm.DB
	once sync.Once
)

func InitDB() (*gorm.DB, error) {
	var err error
	DB, err = gorm.Open(mysql.Open(MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	// if !DB.Migrator().HasTable(&User{}) {
	// 	err = DB.Migrator().CreateTable(&User{})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	//
	// if !DB.Migrator().HasTable(&Video{}) {
	// 	err = DB.Migrator().CreateTable(&Video{})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	//
	// if !DB.Migrator().HasTable(&Comment{}) {
	// 	err = DB.Migrator().CreateTable(&Comment{})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	err = DB.Find(&User{}, &Video{}, &Comment{}).Error
	if err != nil {
		err = DB.AutoMigrate(&User{}, &Video{}, &Comment{})
		if err != nil {
			panic(err)
		}
	}
	return DB, nil
}

func GetDB() (*gorm.DB, error) {
	once.Do(func() {
		_, err := InitDB()
		if err != nil {
			panic(err)
		}
	})
	return DB, nil
}
