/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 17:40
 */

package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDefaultDbEngine() {
	dc := GetConf().Database
	fmt.Println("dc = ", dc)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", dc.Username, dc.Password, dc.Host, dc.Port, dc.Dbname, dc.Config)
	fmt.Println("dsn = ", dsn)
	//dsn := "root:nbroot@tcp(202.204.100.175:3306)/dou_xiao_yin?charset=utf8mb4&parseTime=True&loc=Local"
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
