/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 17:40
 */

package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDefaultDbEngine() {
	dsn := "root:root@tcp(127.0.0.1:3306)/dou_xiao_yin?charset=utf8mb4&parseTime=True&loc=Local"
	// db为单例
	if db == nil {
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	}
}

func GetDefaultDb() *gorm.DB {
	if db == nil {
		InitDefaultDbEngine()
	}
	return db
}
